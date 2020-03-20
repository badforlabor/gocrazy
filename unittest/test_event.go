/**
 * Auth :   liubo
 * Date :   2020/3/20 17:38
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/alg"
	"github.com/stretchr/testify/assert"
	"strconv"
)


type TestEventA struct {
	A string
	Result string
}
func (self *TestEventA) showme(b int) {
	fmt.Println(self.A, b)
	self.Result = self.A + strconv.Itoa(b)
}


func TestEvent() {

	var t = &MyTest{}

	var event = alg.NewEventDispatcher()
	event.Add("1", func() {
		fmt.Println("1")
	})
	event.Invoke("1")

	var a = &TestEventA{A: "aaaaa"}
	var b = &TestEventA{A: "bbbbb"}

	event.Add("a", a.showme)
	event.Add("b", b.showme)

	event.Invoke("a", 1)
	event.Invoke("b", 2)

	assert.Equal(t, a.Result, a.A + strconv.Itoa(1))
	assert.Equal(t, b.Result, b.A + strconv.Itoa(2))

	assert.NotEqual(t, event.Add("a", a.showme), nil)
	assert.NotEqual(t, event.Add("b", b.showme), nil)

	event.Remove("a", a.showme)
	assert.NotEqual(t, event.Invoke("a", 1), nil)

	event.Clear()
	assert.NotEqual(t, event.Invoke("b", 1), nil)
	assert.Equal(t, event.Add("a", a.showme), nil)
}
