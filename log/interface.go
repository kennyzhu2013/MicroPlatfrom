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
	l.Debug(args...)
}
func Info(args ...interface{}) {
	l.Info(args...)
}
func Error(args ...interface{}) {
	l.Error(args...)
}
func Fatal(args ...interface{}) {
	l.Fatal(args...)
}
// Formatted logger
func Debugf(format string, args ...interface{}){
	l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
