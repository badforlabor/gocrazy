/**
 * Auth :   liubo
 * Date :   2019/12/9 14:00
 * Comment:
 */

package crazyio

import (
	"os"
	"path/filepath"
)

// d:/folder1/ -> folder1
// d:/folder1/file.txt -> folder1
// d:/folder1 -> folder1
func GetLastFolder(fullpath string) string {

	for len(fullpath) > 0 && os.IsPathSeparator(fullpath[len(fullpath)-1]) {
		fullpath = fullpath[0 : len(fullpath)-1]
	}

	var idx = 0
	for i:=len(fullpath)-1; i>=0; i-- {
		if os.IsPathSeparator(fullpath[i]) {
			idx = i + 1
			break
		}
	}

	return fullpath[idx:]
}

func Remove(fullpath string)  {
	info, err := os.Stat(fullpath)
	if err == nil {
		if info.IsDir() {
			os.RemoveAll(fullpath)
		} else {
			os.Remove(fullpath)
		}
	}
}

func RemoveAll(wildcardPath string)  {
	files, err := filepath.Glob(wildcardPath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		info, err := os.Stat(f)
		if err == nil {
			if info.IsDir() {
				os.RemoveAll(f)
			} else {
				os.Remove(f)
			}
		}
	}
}