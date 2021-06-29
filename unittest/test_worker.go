/**
 * Auth :   liubo
 * Date :   2021/6/29 17:27
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazytools"
	"sync"
	"time"
)

func testWorker() {

	// 验证是否是单线程的
	{
		var worker = crazytools.NewWorker()
		var useWorker = true
		var locker sync.Mutex
		var locker2 sync.Mutex
		var tt = 2
		var ii = 1
		var wg sync.WaitGroup
		wg.Add(tt * ii)

		var b1 = false
		var b2 = false

		// tt个线程
		for t:=0; t<tt; t++ {
			var threadId = t
			go func() {
				// 每个线程中干ii件事情
				for i:=0; i<ii; i++ {
					var action1 = func() {
						locker.Lock()
						time.Sleep(time.Millisecond * 2000)
						fmt.Println("action1")
						locker2.Lock()
						locker2.Unlock()
						locker.Unlock()
						wg.Done()
						b1 = true
					}
					var action2 = func() {
						locker2.Lock()
						time.Sleep(time.Millisecond * 2000)
						fmt.Println("action2")
						locker.Lock()
						locker.Unlock()
						locker2.Unlock()
						wg.Done()
						b2 = true
					}

					var action = action1
					if threadId % 2 == 0 {
						action = action2
					}

					if useWorker {
						worker.Add(action)
					} else {
						action()
					}
				}
			}()
		}
		if useWorker {
			worker.WaitDone(time.Second)
			if !b1 || !b2 {
				panic("漏掉了某个Action")
			}
		}
		wg.Wait()
	}

	// 验证WaitDone功能
	{
		var worker crazytools.Worker
		worker.Start()

		var last = 0
		var cnt = 0
		var testCnt = 10000
		go func() {
			for i:=0; i<testCnt; i++ {
				var t = i
				worker.Add(func() {
					if t < last {
						panic(t)
					}

					last = t
					cnt++
				})
			}
		}()

		worker.WaitDone(time.Second)

		if cnt != testCnt {
			panic(cnt)
		}
	}

	{
		var worker = crazytools.NewWorker()
		time.Sleep(time.Second * 2)
		if worker.IsDone() {
			panic("没有正确执行")
		}

		var done = false
		worker.Add(func() {
			done = true
		})
		worker.WaitDone(time.Second)
		if !done {
			panic("没有正确执行")
		}
	}

	fmt.Println("test worker done...")
}
