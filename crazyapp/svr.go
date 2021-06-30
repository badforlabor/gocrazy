/**
 * Auth :   liubo
 * Date :   2021/4/7 15:37
 * Comment: 服务框架
 */

package crazyapp

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	//"github.com/emersion/go-autostart"
	"github.com/kardianos/service"
	"github.com/postfinance/single"
	//"os"
	"runtime/debug"
	"time"
)

var svcIsWindowsService func()(bool, error)

type program struct{}
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			fmt.Printf("exception:%s\r\n%s\r\n", err, string(buf))
		}
	}()

	// 代码写在这儿
	mainSvrCallback()
}

func (p *program) Stop(s service.Service) error {
	if mainSvrQuitCallback != nil {
		mainSvrQuitCallback()
	}

	return nil
}

/**
* MAIN函数，程序入口
 */

// 比如以管理员身份，执行如下命令
var install = flag.Bool("install", false, "安装服务")
var remove = flag.Bool("remove", false, "卸载服务")
var delay = flag.Bool("delay", false, "延迟启动")
var failedRestart = flag.Bool("failedRestart", false, "启动失败后，自动重启")

const constDefaultServiceName = "xxx_svr"

var CustomServiceName func() string = nil
var mainSvrCallback func() = nil
var mainSvrQuitCallback func() = nil

func getAppName() string {
	var s = filepath.Base(os.Args[0])
	var n = strings.IndexByte(s, '.')
	if n >= 0 {
		return s[0:n]
	}
	return s
}

func callMain(mainFunc func()) {

	mainSvrCallback = mainFunc

	flag.Parse()

	var defaultServiceName = ""

	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			fmt.Printf("\n ==========================================================================================")
			fmt.Printf("exception:%s\r\n%s\r\n", err, string(buf))
			fmt.Printf("\n ==========================================================================================")
			fmt.Printf("exception:%s\r\n%s\r\n", err, string(buf))
		}
	}()

	if mainSvrCallback == nil {
		panic("invalid main function...")
	}

	if CustomServiceName != nil {
		defaultServiceName = CustomServiceName()
	} else {
		if len(os.Args) > 0 {
			defaultServiceName = getAppName()
		}
	}
	if len(defaultServiceName) == 0 {
		defaultServiceName = constDefaultServiceName
	}

	var sing, e = single.New(defaultServiceName)
	if e != nil {
		// 启动失败
		return
	}

	e = sing.Lock()
	if e != nil {
		return
	}
	defer sing.Unlock()

	fmt.Println("start...", os.Args)

	var opt service.KeyValue = make(map[string]interface{})
	if *delay {
		opt["DelayedAutoStart"] = true
	}
	if *failedRestart {
		opt["OnFailure"] = "restart"
	}

	var svcConfig = &service.Config{
		Name:        defaultServiceName, //服务显示名称
		DisplayName: "",                 //服务名称
		Description: "",                 //服务描述
		Option:      opt,
	}
	//var app = autostart.App{Name:defaultServiceName, DisplayName:defaultServiceName, Exec:[]string{os.Args[0]}}

	fmt.Println("服务名称:", svcConfig.Name)

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	if *install {
		//app.Enable()

		var e = s.Install()
		if e != nil {
			fmt.Println("服务安装失败", e.Error())
		} else {
			fmt.Println("服务安装成功")

			var tryCnt = 3
			for tryCnt > 0 {
				tryCnt--
				time.Sleep(time.Millisecond * 50)
				e = s.Start()
				if e != nil {
					fmt.Println("服务启动失败", e.Error())
					time.Sleep(time.Millisecond * 100)
				} else {
					fmt.Println("服务启动成功")
					break
				}
			}
		}
		return
	}

	if *remove {
		//app.Disable()

		s.Stop()
		time.Sleep(time.Second)
		var e = s.Uninstall()
		fmt.Println("服务卸载成功", e)

		return
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}