package main

import (
	mylog "log/log"
	myfile "log/log/output/file"
	"fmt"
	"os"
	"io"
	"os/signal"
	"path/filepath"
	"syscall"
)

func writeLog(log1 mylog.Log)  {
	for i:=0; i < 10; i++ {
		log1.Infof("this is a info message:%v", i)
	}
}
func main() {
	mylog.FileSize = 1024
	log1 := mylog.NewLog(mylog.WithFields(mylog.Fields{"app":"test"}),
		mylog.WithOutput(myfile.NewOutput(mylog.OutputDir(""),
			mylog.OutputName("server.log"))))
	log1.Options()
	log1.Debug("this is a DEBUG message")
	log1.Info("this is a info message")
	log1.Warn("this is a warn message")
	log1.Error("this is a error message")
	log1.Fatal("this is a fatal message")
	go writeLog(log1)
	// go writeLog(log1)

	fmt.Println( log1.Options().Op )
	// 打开本地文件 读取出全部数据
	fin1, _ := os.Open("server.log")
	fin2, _ := os.Open("/home/wk13802885277/zhuwei/ct-base/go-common/src/testModules/server2.log")
	defer fin1.Close()
	defer fin2.Close()

	fmt.Println("当前sjlin.json1的大小： " )
	fmt.Println(fin1.Seek(0, io.SeekEnd))
	fmt.Println("当前sjlin.json2的大小： " )
	fmt.Println(fin2.Seek(0, io.SeekEnd))

	fileFullPath,_ := filepath.Abs( filepath.Dir( fin1.Name() ) )
	fileFullPath += filepath.Base( fin1.Name() )

	fmt.Println( "fileFullPath is:" + fileFullPath )
	fmt.Println( "fin1.Name() is:" + fin1.Name() )


	fileFullPath,_ = filepath.Abs( filepath.Dir( fin2.Name() ) )
	fileFullPath += filepath.Base( fin2.Name() )

	fmt.Println( "fileFullPath2 is:" + fileFullPath )
	fmt.Println( "fin2.Name() is:" + fin2.Name() )

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-notify

	// fmt.Println(log1.Options().Context)
	// fmt.Println(log1.Options().Fields)
	// fmt.Println(log1.Options().Level)
	// fmt.Println(log1.Options())
	// fmt.Println(log1)
}
