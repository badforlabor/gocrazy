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
	"github.com/badforlabor/gocrazy/crazylog"
)

type MyTest struct {

}
func (self *MyTest) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

var c1 = crazylog.NewCategoryLogDefault("c1", nil)
var c2 = crazylog.NewCategoryLog2("c2", c1)

func main() {
	var action = func() {

		glog.BaseInit()
		glog.Infoln("1\r\n")
		glog.Infoln("2")
		glog.Info("3")
		glog.Warningln("3")

		c1.Infoln("4")
		c1.Infof("5")
		c2.Warnln("6")
		c2.Errorln("7")


		TestEvent()
		testIO()

		testTools()
	}

	if false {
		crazyapp.CallMain(action)
	} else {
		crazyapp.CallNormalMain(action)
	}
}