/**
 * Auth :   liubo
 * Date :   2019/12/25 17:42
 * Comment: 唯一ID
 */

package alg

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"
)

// 根据当前时间进行sha1计算，取前7位。（仿照git）
func GetShortId() string {

	var s = GetHashId()
	return s[:7]
}

func GetHashId() string {

	hash := sha1.New()
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, time.Now().UnixNano())
	hash.Write(buff.Bytes())
	s := fmt.Sprintf("%016x", hash.Sum(nil))

	return s
}