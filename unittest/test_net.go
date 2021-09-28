/**
 * Auth :   liubo
 * Date :   2021/9/27 9:31
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazynet"
)

func testNet() {
	var ip, _ = crazynet.LocalIP()
	fmt.Println("localip", ip.String())
}
