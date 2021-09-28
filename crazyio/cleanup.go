/**
 * Auth :   liubo
 * Date :   2021/9/28 15:16
 * Comment: 清理掉一些文件
 */

package crazyio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

//
//  DeleteFiles
//  @Description: 删除文件夹中的一些文件
//  @param folder，文件夹路径
//  @param regexpStr，文件匹配规则，正则表达式，譬如 `^[a-zA-Z][a-zA-Z0-9]+.*\.log`
//  @param maxAge，保留多久之内的文件，譬如保留7天之内的文件 maxAge = 7 * 24 * time.Hour
//  @return error
//
func DeleteFiles(folder string, regexpStr string, maxAge time.Duration) error {
	if !PathExists(folder) {
		return nil
	}

	files, _ := ioutil.ReadDir(folder)

	var isMatchFile = regexp.MustCompile(regexpStr).MatchString

	var e error = nil

	for _, f := range files {
		if isMatchFile(f.Name()) {
			var e2 = drop(folder, f, maxAge)
			if e2 != nil {
				e = e2
			}
		}
	}

	return e
}
func drop(folder string, f os.FileInfo, maxAge time.Duration) error {
	if time.Since(f.ModTime()) > maxAge {
		return os.Remove(filepath.Join(folder, f.Name()))
	}
	return nil
}
