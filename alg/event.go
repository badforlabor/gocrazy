/**
 * Auth :   liubo
 * Date :   2019/12/26 15:29
 * Comment: 事件派发. 不支持多线程
 */

package alg

import (
	"container/list"
	"errors"
	"reflect"
	"sync"
)

type EventDispatcher struct {
	handler map[string]*list.List
	mux IRWMutex
}

func NewEventDispatcher() *EventDispatcher {

	return &EventDispatcher{
		handler: make(map[string]*list.List),
		mux: NewNilRwLock(),
	}
}
func NewEventDispatcherThread() *EventDispatcher {
	return &EventDispatcher{
		handler: make(map[string]*list.List),
		mux: &sync.RWMutex{},
	}
}

func (self *EventDispatcher) Add(name string, callback interface{}) error {
	self.mux.Lock()
	defer self.mux.Unlock()

	v := self.handler[name]

	if v == nil {
		v = list.New()

		self.handler[name] = v
	}

	var find = false

	for e := v.Front(); e != nil; e = e.Next() {
		if equalFunc(e.Value, callback) {
			find = true
			break
		}
	}

	if !find {
		var v2 = reflect.ValueOf(callback)
		v.PushBack(v2)
		return nil
	}
	return errors.New("duplicate")
}

func (self *EventDispatcher) Invoke(name string, args ...interface{}) error {
	self.mux.RLock()
	defer self.mux.RUnlock()

	var find = false

	if v, ok := self.handler[name]; ok {

		var argsList []reflect.Value
		for _, v := range args {
			argsList = append(argsList, reflect.ValueOf(v))
		}

		for e := v.Front(); e != nil; e = e.Next() {
			c := e.Value.(reflect.Value)
			c.Call(argsList)

			find = true
		}
	}

	if find {
		return nil
	}
	return errors.New("not found")
}

func (self *EventDispatcher) Remove(name string, callback interface{}) error {
	self.mux.Lock()
	defer self.mux.Unlock()

	arr, ok := self.handler[name]
	if !ok {
		return errors.New("not found")
	}

	var next *list.Element
	for e := arr.Front(); e != nil; e = next {
		next = e.Next()

		if equalFunc(e.Value, callback) {
			arr.Remove(e)
		}
	}
	return nil
}

func (self *EventDispatcher) Clear() {
	self.mux.Lock()
	defer self.mux.Unlock()

	self.handler = make(map[string]*list.List)
}

func equalFunc(listValue interface{}, callback interface{}) bool {
	var a = listValue.(reflect.Value)
	return a.Pointer() == reflect.ValueOf(callback).Pointer()
}