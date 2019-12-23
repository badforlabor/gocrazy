/**
 * Auth :   liubo
 * Date :   2019/12/3 15:51
 * Comment: 查找
 */

package alg

// 标记是int
func IntCompare(a int) func(d interface{})bool {
	var function = func(d interface{})bool {
		var aa = d.(int)
		return a == aa
	}
	return function
}

func Find(dataList interface{}, compare func(d interface{})bool) int {
	if !isSlice(dataList) {
		return -2
	}
	return findList(ToSlice(dataList), compare)
}

func Contain(dataList interface{}, compare func(d interface{})bool) bool {
	return Find(dataList, compare) != -1
}

func findList(dataList []interface{}, compare func(d interface{})bool) int {
	for i:=0; i<len(dataList); i++ {
		if compare(dataList[i]) {
			return i
		}
	}
	return -1
}
