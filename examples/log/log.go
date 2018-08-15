package main

import (
	"time"

	"github.com/kennyzhu/go-os/log"
)

func main() {
	logger := log.NewLog(
		log.WithLevel(log.InfoLevel),
		log.WithFields(log.Fields{
			"logger": "os",
		}),
		log.WithOutput(
			log.NewOutput(log.OutputName("/dev/stdout")), //output dir...
		),
	)

	for i := 0; i < 100; i++ {
		logger.Info("This is a log message")
		time.Sleep(time.Second)
	}
}
