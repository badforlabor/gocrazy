/**
 * Auth :   liubo
 * Date :   2020/7/7 9:59
 * Comment: 异步等待，是一种阻塞等待. todo，清理
 */

package crazytools


import (
	"sync"
	"sync/atomic"
	"time"
)

var _waitId uint64

var mapChan = make(map[uint64]chan interface{})
var mutexMapChan sync.RWMutex

func GetWaitId() uint64 {
	var id = atomic.AddUint64(&_waitId, 1)
	return id
}

// 返回"是否通知成功"
func Notify(id uint64, userData interface{}) bool {
	mutexMapChan.RLock()
	var v, ok = mapChan[id]
	mutexMapChan.RUnlock()

	if ok {
		removeKey(id)

		v <- userData
	}
	return ok
}

func HasId(id uint64) bool {
	mutexMapChan.RLock()
	var _, ok = mapChan[id]
	mutexMapChan.RUnlock()
	return ok
}

func removeKey(id uint64) {
	mutexMapChan.Lock()
	defer mutexMapChan.Unlock()

	delete(mapChan, id)
}

// 返回“是否超时”
func Wait(id uint64, waitMaxTime time.Duration) (bool, interface{}) {
	var c = make(chan interface{})
	mutexMapChan.Lock()
	mapChan[id] = c
	mutexMapChan.Unlock()

	var timeout = true
	var userData interface{} = nil
	
	select {
	case d := <-c :
		timeout = false
		userData = d
		break

	case <-time.After(waitMaxTime):
		removeKey(id)
		break
	}

	return timeout, userData
}

