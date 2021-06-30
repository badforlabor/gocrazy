/**
 * Auth :   liubo
 * Date :   2020/3/20 17:43
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"github.com/badforlabor/gocrazy/crazyapp"
)

type MyTest struct {

}
func (self *MyTest) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func main() {
	crazyapp.CallMain(func() {

		glog.BaseInit()
		glog.Infoln("1\r\n")
		glog.Infoln("2")
		glog.Info("3")
		glog.Warningln("3")


		TestEvent()
		testIO()

		testTools()
	})
}