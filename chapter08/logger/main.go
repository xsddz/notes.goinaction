// 这个示例程序展示如何创建定制的日志记录器
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace 记录所有日志
	Trace *log.Logger
	// Info 重要的信息
	Info *log.Logger
	// Warning 需要注意的信息
	Warning *log.Logger
	// Error 非常严重的问题
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	// MultiWriter 函数是一个变参函数，可以接受任意个实现了 io.Writer 接口的值。
	// 这个函数会返回一个 io.Writer 值，这个值会把所有传入的 io.Writer 的值绑在一起。
	// 当对这个返回值进行写入时，会向所有绑在一起的 io.Writer 值做写入。
	// 这让类似 log.New 这样的函数可以同时向多个 Writer 做输出。
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
