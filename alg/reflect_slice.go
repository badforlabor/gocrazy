/**
 * Auth :   liubo
 * Date :   2019/12/3 16:10
 * Comment: Slice的扩展函数，以Slice开头
 */

package alg

import "reflect"

func isSlice(dataList interface{}) bool {
	s := reflect.ValueOf(dataList)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	if s.Kind() == reflect.Slice {
		return true
	}
	return false
}

func reflectSlice(dataList interface{}) reflect.Value {
	s := reflect.ValueOf(dataList)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	if s.Kind() != reflect.Slice {
		panic("not slice")
	}

	return s
}

// 拆箱
func ToSlice(dataList interface{}) []interface{} {
	s := reflectSlice(dataList)

	ret := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// 再次拆箱
func SetSliceValue(src []interface{}, dst interface{}) {
	s := reflectSlice(dst)

	if s.Len() != len(src) {
		panic("not equal")
	}

	for i:=0; i<s.Len(); i++ {
		s.Index(i).Set(reflect.ValueOf(src[i]))
	}
}

// 查找
func SliceFind(dataList interface{}, compare func(d interface{})bool) int {
	return Find(dataList, compare)
}
func SliceContain(dataList interface{}, compare func(d interface{})bool) bool {
	return Find(dataList, compare) != -1
}

// 删除
func SliceRemoveAt(dataList interface{}, idx int) {
	mustPtr(dataList)

	if idx < 0 {
		panic("invalid index")
	}

	s := reflectSlice(dataList)
	if s.Len() == 0 {
		return
	}

	for i:=idx; i<s.Len()-1; i++ {
		s.Index(i).Set(s.Index(i+1))
	}

	//s.SetLen(s.Len() - 1)
	s.SetLen(s.Len() - 1)
}
func SliceRemoveOne(dataList interface{}, compare func(d interface{})bool) {
	idx := Find(dataList, compare)
	if idx >= 0 {
		SliceRemoveAt(dataList, idx)
	}
}

// 增加
func SliceAddOne(dataList, one interface{}) {
	mustPtr(dataList)

	SliceInsertOne(dataList, -1, one)
}
func SliceAddUnique(dataList, data interface{}, compare func(d interface{})bool) {
	if Contain(dataList, compare) {
		return
	}
	SliceAddOne(dataList, data)
}
func SliceInsertOne(dataList interface{}, idx int, one interface{}) {
	mustPtr(dataList)

	if idx < 0 && idx != -1 {
		panic("invalid index")
	}

	s := reflectSlice(dataList)

	if s.Len()+1 > s.Cap() {
		dst := reflect.MakeSlice(s.Type(), s.Len() + 1, s.Cap() * 2)
		reflect.Copy(dst, s)
		s.Set(dst)
	} else {
		s.SetLen(s.Len() + 1)
	}

	if idx == -1{
		idx = s.Len() - 1
	} else {
		for i:=s.Len()-1; i>idx; i-- {
			s.Index(i).Set(s.Index(i-1))
		}
	}
	s.Index(idx).Set(reflect.ValueOf(one))
}

// 反转过来
func SliceInverse(datalist interface{}) {
	//mustPtr(datalist)
	s := reflectSlice(datalist)
	mid := s.Len() / 2
	for i:=0; i<mid; i++ {
		j := s.Len() - 1 - i

		x, y := s.Index(i).Interface(), s.Index(j).Interface()
		s.Index(i).Set(reflect.ValueOf(y))
		s.Index(j).Set(reflect.ValueOf(x))
	}
}