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
	"io/ioutil"
	"os"
	"reflect"
)
import _ "github.com/badforlabor/gocrazy/crazylog"
import _ "github.com/badforlabor/gocrazy/crazyos"
import _ "github.com/badforlabor/gocrazy/crazytools"

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

	{
		os.MkdirAll("c:/1/2", os.ModePerm)
		os.MkdirAll("c:/1/2/3", os.ModePerm)
		os.MkdirAll("c:/1/2/4", os.ModePerm)
		os.MkdirAll("c:/1/2/5", os.ModePerm)
		ioutil.WriteFile("c:/1/2/3/1.txt", []byte("1234"), os.ModePerm)
		var e = crazyio.MovePath("c:\\1\\2", "d:/1/2")
		if e == nil {
			fmt.Println("移动文件成功")
		} else {
			fmt.Println("移动文件失败", e.Error())
		}

		if crazyio.PathExists("c:/1/2/4") {
			fmt.Println("自动删除文件失败")
		}
	}

	{
		var folders = crazyio.GetFolders("D:\\Program Files (x86)", ".*")
		fmt.Println(folders)
	}
	{
		var folders = crazyio.GetFolders("./", ".*")
		fmt.Println(folders)
	}
}
