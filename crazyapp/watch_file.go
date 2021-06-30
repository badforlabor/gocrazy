/**
 * Auth :   liubo
 * Date :   2021/6/30 11:36
 * Comment: 监控文件变化，带扩展
 */

package crazyapp

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"github.com/fsnotify/fsnotify"
)

func WatchMonify(filename string, callback func()) *fsnotify.Watcher {
	// 监控配置文件变化
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		glog.Warningf("创建watcher失败:%v", filename)
		return nil
	}

	OnQuit(func() {
		watcher.Close()
	})
	go func() {
		watcherEvent(watcher, func(file string) {
			if file == filename {
				callback()
			}
		})
	}()
	err = watcher.Add(filename)
	if err != nil {
		glog.Errorf("监视文件失败！err=%v, filename=%v", err.Error(), filename)
	} else {
		glog.Infof("监视文件:%v", filename)
	}
	return watcher
}

func watcherEvent(watcher *fsnotify.Watcher, onModify func(filename string)) {

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("modified file:", event.Name)
				onModify(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("error:", err)
		}
	}
}

