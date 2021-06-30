/**
 * Auth :   liubo
 * Date :   2021/6/30 11:54
 * Comment: 封装一般的app
 */

package crazyapp

import (
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"os"
	"os/signal"
	"runtime/debug"
)

// 主进程。用来支持“单线程”工作
var MainThread *Worker

// 服务类型的app
func CallMain(mainFunc func()) {

	mainSvrQuitCallback = QuitAction

	callMain(func() {
		BaseInit()
		mainFunc()
	})
}

// 普通app
func CallNormalMain(mainFunc func()) {
	BaseInit()
	mainFunc()
	WaitQuit()
}

func WaitQuit() {

	// 监听终止
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	glog.Info("等待信号")
	<-c
	glog.Info("收到信号")

	QuitAction()
}
func QuitAction() {

	for _, v := range onQuitCallback {
		RunCallback(v)
	}

	StopCron()
	MainThread.Stop()

	glog.Flush()
}

var onQuitCallback []func()

func OnQuit(callback func()) {
	onQuitCallback = append(onQuitCallback, callback)
}
func RunCallback(callback func()) {
	if callback == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			glog.Warningln("exception:%s\r\n%s\r\n", err, string(buf))
		}
	}()

	callback()
}


func BaseInit() {
	glog.BaseInit()

	MainThread = NewWorker()

}