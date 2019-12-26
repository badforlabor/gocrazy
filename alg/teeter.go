/**
 * Auth :   liubo
 * Date :   2019/12/26 11:58
 * Comment: 同步阻塞执行一系列异步函数
 */

package alg

import (
	"fmt"
	"time"
)

type Teeter struct {
	t chan bool
	timeout bool
}
func (self *Teeter) Sync(maxWaitTime time.Duration, callback func()) {
	self.Begin(maxWaitTime, func() {
		callback()
		self.End()
	})
}
func (self *Teeter) IsTimeout() bool {
	return self.timeout
}
func (self *Teeter) Begin(maxWaitTime time.Duration, callback func()) {
	self.timeout = false
	self.t = make(chan bool)

	go func() {
		callback()
	}()

	select {
	case <-self.t:
		{
			//fmt.Println("autowait done")
		}
	case <-time.After(maxWaitTime):
		self.timeout = true
		//fmt.Println("autowait timeout")
	}
}
func (self *Teeter) End() {
	self.t <- true
}

func testAutoWait() {
	var a Teeter
	a.Sync(3 * time.Second, func() {
		c := time.After(time.Second * 2)
		<-c

		if a.IsTimeout() {
			return
		}
		fmt.Println("step1")

		<-time.After(time.Second * 2)

		if a.IsTimeout() {
			return
		}
		fmt.Println("step2")
	})

	<-time.After(time.Second * 7)
}
