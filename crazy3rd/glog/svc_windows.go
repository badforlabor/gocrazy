// +build windows

/**
 * Auth :   liubo
 * Date :   2021/4/14 16:59
 * Comment:
 */

package glog


import (
	"golang.org/x/sys/windows/svc"
)

func init() {
	svcIsWindowsService = svc.IsWindowsService
}
