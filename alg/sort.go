/**
 * Auth :   liubo
 * Date :   2019/12/3 11:10
 * Comment: 算法：排序
 */

package alg

import "sort"

// 通用排序
func Sort(dataList interface{}, compare func(left interface{}, right interface{}) bool ) {
	sortAscending(true, dataList, compare)
}
// 反向排序
func SortInverse(dataList interface{}, compare func(left interface{}, right interface{}) bool ) {
	sortAscending(false, dataList, compare)
}
func sortAscending(ascending bool, dataList interface{}, compare func(left interface{}, right interface{}) bool) {
	data := commonList{dataList:ToSlice(dataList), compare:compare, ascending:ascending}
	sort.Sort(&data)
	SetSliceValue(data.dataList, dataList)
}
func SortInt() func(left interface{}, right interface{}) bool {
	var compare = func(left interface{}, right interface{}) bool {
		a := left.(int)
		b := right.(int)
		return a < b
	}
	return compare
}

type commonList struct {
	dataList []interface{}
	compare func(left interface{}, right interface{}) bool
	ascending bool
}
func (self *commonList) Len() int {
	return len(self.dataList)
}
func (self *commonList) Less(i, j int) bool {
	return self.ascending == self.compare(self.dataList[i], self.dataList[j])
}
func (self *commonList) Swap(i, j int) {
	var tmp = self.dataList[i]
	self.dataList[i] = self.dataList[j]
	self.dataList[j] = tmp
}
