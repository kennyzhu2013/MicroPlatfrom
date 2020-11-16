/*
@Time : 2018/8/16 12:46 
@Author : kenny zhu
@File : interface.go
@Software: GoLand
@Others:
*/
package log

//logger
var l Log

func InitLogger(opts ...Option)  {
	l = newOS(opts...)
}

//outer .interface...
func Debug(args ...interface{}) {
	Debug(args...)
}
func Info(args ...interface{}) {
	Info(args...)
}
func Error(args ...interface{}) {
	Error(args...)
}
func Fatal(args ...interface{}) {
	Fatal(args...)
}
// Formatted logger
func Debugf(format string, args ...interface{}){
	Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	Infof(format, args...)
}
func Errorf(format string, args ...interface{}) {
	Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	Fatalf(format, args...)
}
