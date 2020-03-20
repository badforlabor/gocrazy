/**
 * Auth :   liubo
 * Date :   2019/12/3 9:26
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/badforlabor/gocrazy/alg"
	"github.com/badforlabor/gocrazy/crazyio"
	"reflect"
)
import _ "github.com/badforlabor/gocrazy/crazylog"
import _ "github.com/badforlabor/gocrazy/crazyos"

func testIO() {
	crazyio.AppendFile("test.txt", "123")

	{
		var datalist = []int{0, 1, 2, 3}
		var idx = alg.Find(datalist, alg.IntCompare(3))
		fmt.Println(idx, idx == 3)
	}

	{
		var datalist = []int{3,2,1,0}
		alg.Sort(datalist, alg.SortInt())
		fmt.Println(datalist)
	}

	{
		var datalist = []int{0, 1, 2, 3}
		v1 := reflect.ValueOf(datalist)
		v2 := reflect.ValueOf(&datalist)
		fmt.Println(v1.Kind(), v2.Kind(), reflect.Ptr)
	}

	var t = &MyTest{}
	{
		TestSlice1(t)
		TestSlice2(t)
		TestClone(t)
	}
}
