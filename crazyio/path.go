/**
 * Auth :   liubo
 * Date :   2019/12/9 14:00
 * Comment:
 */

package crazyio

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 格式化路径（转化成操作系统格式的路径）
func FormatPath(fullPath string) string {
	fullPath = filepath.Clean(fullPath)
	//var sep = string(os.PathSeparator)
	//return strings.Join(strings.Split(fullPath, sep), sep)
	return fullPath
}

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

// 获取exe的名字
func GetExeName(fullpath string) string {
	var s = filepath.Base(fullpath)
	var n = strings.IndexByte(s, '.')
	if n >= 0 {
		return s[0:n]
	}
	return s
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
func Mkdir(fullpath string) error {
	return os.MkdirAll(fullpath, os.ModePerm)
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

// 拷贝文件，或者文件夹
func CopyPath(srcFullpath, dstFullpath string) error {
	var f, e = os.Stat(srcFullpath)
	if e != nil {
		return e
	}
	if f.IsDir() {
		e = copyDir(srcFullpath, dstFullpath)
	} else {
		_, e = copyFile(srcFullpath, dstFullpath)
	}
	return e
}

// 移动文件夹
func MovePath(srcFullpath, dstFullpath string) error {
	var e = os.Rename(srcFullpath, dstFullpath)
	if e != nil {
		// 手动移动
		var f, e = os.Stat(srcFullpath)
		if e == nil {
			if f.IsDir() {
				filepath.Walk(srcFullpath, func(path string, info os.FileInfo, err error) error {
					if !info.IsDir() {
						var _, e2 = copyFile(path, filepath.Join(dstFullpath, path[len(srcFullpath):]))
						if e2 == nil {
							os.Remove(path)
						} else {
							e = e2
						}
					} else if len(path) > len(srcFullpath) {
						Mkdir(filepath.Join(dstFullpath, path[len(srcFullpath):]))
					}
					return nil
				})
			}
		} else {
			_, e = copyFile(srcFullpath, dstFullpath)
			if e == nil {
				os.Remove(srcFullpath)
			}
		}
	}
	return e
}

// 获取文件夹（只获取第一层）
func GetFolders(pathName string, regExpStr string) []string{
	var r, _ = regexp.Compile(regExpStr)

	if len(regExpStr) == 0 {
		r = nil
	}

	var ret []string
	filepath.Walk(pathName, func(path string, info os.FileInfo, err error) error {

		// 自身
		if len(path) == len(pathName) {
			return nil
		}

		if info.IsDir() {
			if r == nil || r.MatchString(info.Name()) {
				ret = append(ret, path)
			}
			return filepath.SkipDir
		}
		return nil
	})
	return ret
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func copyFile(src, dst string) (w int64, err error) {
	w = 0
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstSlices := strings.Split(dst, string(os.PathSeparator))
	dstSlicesLen := len(dstSlices)
	destDir := ""
	for i := 0; i < dstSlicesLen-1; i++ {
		destDir = destDir + dstSlices[i] + string(os.PathSeparator)
	}
	b, err := pathExists(destDir)
	if b == false {
		err = os.MkdirAll(destDir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			return
		}
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func copyDir(src string, dest string) error {
	srcOriginal := src
	// 否则使用'ioutil.ReadDir'来自行逐层遍历!
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		filePart, _ := filepath.Rel(srcOriginal, src)
		destNew := filepath.Join(dest, filePart)

		// 注意:这里会walk所有得文件!(而不仅仅是当前层级得)
		if !f.IsDir() {
			_, err := copyFile(src, destNew)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
