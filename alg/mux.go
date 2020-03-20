/**
 * Auth :   liubo
 * Date :   2020/3/20 17:35
 * Comment:
 */

package alg

import "sync"

type IRWMutex interface {
	sync.Locker
	RLock()
	RUnlock()
}
func NewNilRwLock() IRWMutex {
	return &nilRwLock{}
}
func NewNilLock() sync.Locker {
	return &nilLock{}
}

type nilLock struct {

}

func (n *nilLock) Lock() {

}

func (n *nilLock) Unlock() {

}

type nilRwLock struct {

}

func (self *nilRwLock) Unlock() {

}

func (self *nilRwLock) RLock() {

}

func (self *nilRwLock) RUnlock() {

}

func (self *nilRwLock) Lock() {

}