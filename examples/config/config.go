package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/micro/go-os/config"
	"github.com/micro/go-os/config/source/file"
)

var (
	configFile = filepath.Join(os.TempDir(), "config.example")
)

func writeFile(i int) error {
	return ioutil.WriteFile(configFile, []byte(fmt.Sprintf(`{"key": "value-%d"}`, i)), 0600)
}

func removeFile() error {
	return os.Remove(configFile)
}

func watch(c config.Config, key ...string) {
	fmt.Println("Subscribing to changes for", key)

	w, err := c.Watch(key...)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			v, err := w.Next()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Received value for key:", v.String("default"))
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Writing change to file")
		writeFile(i + 1)
		time.Sleep(time.Second)
	}

	fmt.Println("Stopping subscriber")
	w.Stop()
}

func main() {
	flag.Parse()

	// Write our first entry
	if err := writeFile(0); err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(configFile)

	// Create a config instance
	config := config.NewConfig(
		// aggressive config polling
		config.PollInterval(time.Millisecond*500),
		// use file as a config source
		config.WithSource(file.NewSource(config.SourceName(configFile))),
	)

	defer config.Close()

	// lets read the value while editing it a number of times
	for i := 0; i < 10; i++ {
		val := config.Get("key").String("default")
		fmt.Println("Got ", val)
		writeFile(i + 1)
		time.Sleep(time.Second)
	}

	// watch key that exists
	watch(config, "key")

	// watch key that does not exist
	watch(config, "foo")

	fmt.Println("Stopping config runner")
}
