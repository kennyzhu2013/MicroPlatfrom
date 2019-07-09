/*
@Time : 2018/8/16 12:46 
@Author : kenny zhu
@File : interface.go
@Software: GoLand
@Others: Default operations.
*/
package log

// logger
var defaultLog Log

func InitLogger(opts ...Option)  {
	defaultLog = newOS(opts...)
}

// outer .interface...
func Debug(args ...interface{}) {
	defaultLog.Debug(args...)
}
func Info(args ...interface{}) {
	defaultLog.Info(args...)
}
func Error(args ...interface{}) {
	defaultLog.Error(args...)
}

func Warn(args ...interface{}) {
	defaultLog.Warn(args...)
}

func Fatal(args ...interface{}) {
	defaultLog.Fatal(args...)
}
// Formatted logger
func Debugf(format string, args ...interface{}){
	defaultLog.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	defaultLog.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLog.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLog.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	defaultLog.Fatalf(format, args...)
}
