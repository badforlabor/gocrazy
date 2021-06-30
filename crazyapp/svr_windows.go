/**
 * Auth :   liubo
 * Date :   2021/6/30 17:07
 * Comment:
 */

package crazyapp

import "golang.org/x/sys/windows/svc"

func init() {
	svcIsWindowsService = svc.IsWindowsService
}
