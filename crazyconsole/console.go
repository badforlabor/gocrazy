/**
 * Auth :   liubo
 * Date :   2020/9/25 14:13
 * Comment:
 */

package crazyconsole

import "fmt"

func Pause() {
	fmt.Println("按‘回车键’继续...")
	fmt.Scanln()
}