package file

/*
	File output can be used to send log messages to files.
	Remember that even stdout/stderr can be used as files
	with /dev/stdout /dev/stderr
*/

import (
	log2 "github.com/kennyzhu/go-os/src/log"
)

func NewOutput(opts ...log2.OutputOption) log2.Output {
	return log2.NewOutput(opts...)
}
