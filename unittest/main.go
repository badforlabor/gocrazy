/**
 * Auth :   liubo
 * Date :   2020/3/20 17:43
 * Comment:
 */

package main

import "fmt"

type MyTest struct {

}
func (self *MyTest) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func main() {
	TestEvent()
	testIO()

	testTools()
}