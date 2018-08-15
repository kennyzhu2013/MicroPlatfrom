package file

/*
	File output can be used to send log messages to files.
	Remember that even stdout/stderr can be used as files
	with /dev/stdout /dev/stderr
*/

import (
	"github.com/micro/go-os/log"
)

func NewOutput(opts ...log.OutputOption) log.Output {
	return log.NewOutput(opts...)
}
