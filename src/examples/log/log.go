package main

import (
	log2 "github.com/kennyzhu/go-os/src/log"
	"time"
)

func main() {
	logger := log2.NewLog(
		log2.WithLevel(log2.InfoLevel),
		log2.WithFields(log2.Fields{
			"logger": "os",
		}),
		log2.WithOutput(
			log2.NewOutput(log2.OutputName("/dev/stdout")), //output dir...
		),
	)

	for i := 0; i < 100; i++ {
		logger.Info("This is a log message")
		time.Sleep(time.Second)
	}
}
