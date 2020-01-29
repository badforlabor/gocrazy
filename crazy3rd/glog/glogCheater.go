package glog

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dir string = ""

func BaseInit() {
	dir := getCurPath()
	dir = strings.Join([]string{dir, "logs"}, string(os.PathSeparator))
	os.Mkdir(dir, os.ModeDir)

	Cheat(dir)

	MaxSize = 1024 * 1024 * 16
}

func Cheat(tmpdir string) {
	dir = tmpdir
	// 设置输出路径
	logDir = &dir
}

func getCurPath() string {
	fmt.Println(os.Args[0])
	file, _ := exec.LookPath(os.Args[0])

	if len(file) == 0 {
		file = os.Args[0]
	}

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	//将全路径用\\分割，得到4段，①E: ②golang ③test ④a.exe
	splitstring := strings.Split(path, string(os.PathSeparator))

	splitstring = splitstring[:len(splitstring) - 1]

	var rst = strings.Join(splitstring, string(os.PathSeparator))

	return rst
}