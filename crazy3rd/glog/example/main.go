package main

import (
	"flag"
	"gocrazy/crazy3rd/glog"
	"time"
)

func main() {

	flag.Parse()

	glog.BaseInit()
	defer glog.Flush()

	glog.Infoln("123")
	glog.Errorln("e123")
	glog.Warningln("w123")

	// 多线程写日志
	func1 := func(tag string){
		for i:=0; i< 100000; i++ {
			glog.Infoln(tag, "123")
			glog.Errorln(tag, "e123")
			glog.Warningln(tag, "w123")
		}
	}
	go func1("func1")
	go func1("func2")
	go func1("func3")
	go func1("func4")
	go func1("func5")

	// 等待协程函数完毕
	time.Sleep(10 * time.Second)
}
