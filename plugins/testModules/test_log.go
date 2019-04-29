package main

import (
	mylog "log/log"
	myfile "log/log/output/file"
	"fmt"
	"os"
	"io"
)
func main() {

	log1 := mylog.NewLog(mylog.WithOutput(myfile.NewOutput(mylog.OutputName("server.log"))))
	log1.Options()
	log1.Debug("this is a DEBUG message")
	log1.Info("this is a info message")
	log1.Warn("this is a warn message")
	log1.Error("this is a error message")
	log1.Fatal("this is a fatal message")
	for i:=0; i < 100; i++ {
		log1.Infof("this is a info message:%v", i)
	}
	fmt.Println( log1.Options().Op )
	// 打开本地文件 读取出全部数据
	fin1, _ := os.Open("server.log")
	fin2, _ := os.Open("server2.log")
	defer fin1.Close()
	defer fin2.Close()

	fmt.Println("当前sjlin.json1的大小： " )
	fmt.Println(fin1.Seek(0, io.SeekEnd))
	fmt.Println("当前sjlin.json2的大小： " )
	fmt.Println(fin2.Seek(0, io.SeekEnd))
	// fmt.Println(log1.Options().Context)
	// fmt.Println(log1.Options().Fields)
	// fmt.Println(log1.Options().Level)
	// fmt.Println(log1.Options())
	// fmt.Println(log1)
}