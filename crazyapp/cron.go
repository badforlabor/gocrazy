/**
 * Auth :   liubo
 * Date :   2021/6/29 14:24
 * Comment: 定时功能
 */

package crazyapp

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

var globalCron *cron.Cron

// cron规则：秒，分，时，天，月，周几
func StartCron(dateTime string, callback func())  {
	if globalCron == nil {
		globalCron = cron.New(cron.WithSeconds())
		globalCron.Start()
	}

	if callback == nil || len(dateTime) == 0 {
		return
	}

	var _, e = globalCron.AddFunc(dateTime, callback)
	if e != nil {
		panic(e.Error())
	}

	// 测试，每隔几秒执行一次
	if false {
		_, e = globalCron.AddFunc("*/5 * * * * *", cronTestJob)
		if e != nil {
			panic(e.Error())
		}
	}

}
func StopCron() {
	if globalCron != nil {
		globalCron.Stop()
	}
}
func cronTestJob() {
	fmt.Println("daily job")
}

