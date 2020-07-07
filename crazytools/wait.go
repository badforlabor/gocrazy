/**
 * Auth :   liubo
 * Date :   2020/7/7 9:59
 * Comment: 异步等待，是一种阻塞等待
 */

package crazytools


import (
	"sync/atomic"
	"time"
)

var _waitId uint64

var mapChan = make(map[uint64]chan bool)

func GetWaitId() uint64 {
	var id = atomic.AddUint64(&_waitId, 1)
	return id
}

// 返回"是否通知成功"
func Notify(id uint64) bool {
	var v, ok = mapChan[id]
	if ok {
		v <- true
	}
	return ok
}

// 返回“是否超时”
func Wait(id uint64, waitMaxTime time.Duration) bool {
	var c = make(chan bool)
	mapChan[id] = c

	var timeout = true

	select {
	case <-c :
		timeout = false
		break

	case <-time.After(waitMaxTime):
		break
	}

	return timeout
}

