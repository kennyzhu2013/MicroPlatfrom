/*
@Time : 2019/5/22 11:57
@Author : kenny zhu
@File : debug
@Software: GoLand
@Others:
*/
package logdebug

import (
	"log/log"
	"runtime"
	"runtime/debug"
)
const FullStackSize = 1 << 12

func LogAllStacks()  {
	// debug.PrintStack(), print to tty.
	log.Errorf("local stack is: %s", debug.Stack() )
	buf := make([]byte, FullStackSize)
	n := runtime.Stack(buf, true)
	log.Errorf("the full stack is: \n=======================================%v\n==========================", string(buf[:n]) )
}

func LogLocalStacks()  {
	// debug.PrintStack(), print to tty.
	log.Errorf("local stack is: \n==========================%s\n==========================", debug.Stack())
}