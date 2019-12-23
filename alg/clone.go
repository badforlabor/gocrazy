/**
 * Auth :   liubo
 * Date :   2019/12/7 20:12
 * Comment: 克隆
 */

package alg

import "reflect"

// 克隆
func Clone(dst, src interface{}) {
	mustPtr(dst)

	if isSlice(src) {
		cloneSlice(dst, src)
		return
	}

	panic("not support clone.")
}
func cloneSlice(dst, src interface{}) {
	s := reflectSlice(src)
	dstValue := reflect.ValueOf(dst).Elem()

	dstValue.Set(reflect.MakeSlice(s.Type(), s.Len(), s.Cap()))
	reflect.Copy(dstValue, s)
}