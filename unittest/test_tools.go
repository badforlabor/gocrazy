/**
 * Auth :   liubo
 * Date :   2020/7/7 10:01
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazytools"
	"time"
)

func testTools() {

	testWait()
	testWait2()

}

func testWait() {

	var id = crazytools.GetWaitId()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("wait 1")
		crazytools.Notify(id)
	}()
	var timeout = crazytools.Wait(id, time.Second * 2)
	fmt.Println("wait 2, timeout=", timeout)

}

func testWait2() {

	var id = crazytools.GetWaitId()

	// 这种，会直接导致超时
	fmt.Println("wait 1")
	var succ = crazytools.Notify(id)
	fmt.Println("notify succ:", succ)

	var timeout = crazytools.Wait(id, time.Second * 2)
	fmt.Println("wait 2, timeout=", timeout)

}