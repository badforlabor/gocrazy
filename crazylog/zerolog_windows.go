/**
 * Auth :   liubo
 * Date :   2021/9/28 16:19
 * Comment:
 */

package crazylog

import "golang.org/x/sys/windows/svc"

func init() {
	svcIsWindowsService = svc.IsWindowsService
}



