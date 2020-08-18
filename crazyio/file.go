/**
 * Auth :   liubo
 * Date :   2019/12/3 9:19
 * Comment: 文件
 */

package crazyio

import (
	"github.com/badforlabor/gocrazy/crazylog"
	"os"
	"path/filepath"
	"strings"
)

// 追加内容
func AppendFile(filename string, text string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		crazylog.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		crazylog.Println(err)
	}
}

// 文件名
func GetFileName(fullpath string) string {
	var path = filepath.SplitList(fullpath)
	if len(path) > 1 {
		path = path[len(path)-1:]
	}

	if len(path) == 1 {
		return path[0]
	}

	return fullpath
}

// 文件夹名字
func GetFolderName(fullpath string) string {
	var path = filepath.SplitList(fullpath)
	if len(path) > 0 {
		if strings.Contains(path[len(path)-1], ".") {
			path = path[0:len(path)-1]
		}
	}
	return filepath.Join(path...)
}
