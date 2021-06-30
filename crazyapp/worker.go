/**
 * Auth :   liubo
 * Date :   2021/6/29 17:00
 * Comment: 工作队列，一个Worker只能单线程工作
 */

package crazyapp

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	OnPanic func(interface{}, *Worker)

	jobList []interface{}
	jobListMutex sync.Mutex

	stop bool
	endSignal sync.WaitGroup

	RemainList []interface{}

	waitTime        time.Duration
	waitTimeRuntime time.Duration
}

// 创建新的，并自动启动
func NewWorker() *Worker {
	var w = Worker{}
	w.Start()
	return &w
}

// 启动
func(self *Worker) Start() {
	self.endSignal.Add(1)
	go func() {

		var done = false
		for !self.stop && !done {

			// 交换
			self.jobListMutex.Lock()
			var pendingList = make([]interface{}, 0)
			pendingList, self.jobList = self.jobList, pendingList
			self.jobListMutex.Unlock()

			if len(pendingList) > 0 {

				for _, v := range pendingList {
					if self.stop {
						// 记录下未完成的
						self.RemainList = append(self.RemainList, v)
					} else {
						// 执行
						switch callback := v.(type) {
						case func():
							self.do(callback)

						case nil:

						default:
							panic(fmt.Sprintf("unknown type:%v", callback))
						}
					}
				}
				self.waitTimeRuntime = 0

			} else {

				// 没有活儿时，休眠
				var t = time.Millisecond
				time.Sleep(t)
				self.waitTimeRuntime += t

				if self.waitTimeRuntime > self.waitTime && self.waitTime > 0 {
					done = true
				}
			}

		}

		// 记录下未完成的
		self.jobListMutex.Lock()
		self.RemainList = append(self.RemainList, self.jobList...)
		self.jobListMutex.Unlock()

		// 记录下未完成的
		self.endSignal.Done()
	}()
}

// 添加内容
func(self *Worker) Add(callback interface{}) {
	if self.stop {
		return
	}

	self.jobListMutex.Lock()
	self.jobList = append(self.jobList, callback)
	self.jobListMutex.Unlock()
}

// 停止
func(self *Worker) Stop() {
	self.stop = true
	self.endSignal.Wait()
}

// 一段时间没有数据，那么就结束工作
func(self *Worker) WaitDone(waitTime time.Duration) {
	self.waitTime = waitTime

	self.endSignal.Wait()
	self.stop = true
}

func(self *Worker) wait() {
	self.WaitDone(time.Hour * 24 * 365 * 10)
}

func(self *Worker) do(callback func()) {
	if self.OnPanic != nil {
		defer func() {
			if err := recover(); err != nil {
				self.OnPanic(err, self)
			}
		}()
	}

	callback()
}