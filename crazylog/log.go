/**
 * Auth :   liubo
 * Date :   2019/12/3 9:23
 * Comment: 库中应用到的日志
 */

package crazylog

import (
	"fmt"
)

func SetPrintlnCallback() {

}
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
