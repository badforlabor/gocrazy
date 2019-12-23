/**
 * Auth :   liubo
 * Date :   2019/12/7 19:19
 * Comment:
 */

package alg

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSort1(t *testing.T) {
	{
		var datalist = []int{5, 0,1,2,3,4}
		Sort(datalist, SortInt())
		for i, v := range datalist {
			assert.Equal(t, i, v)
		}

		SortInverse(datalist, SortInt())
		for i, v := range datalist {
			assert.Equal(t, len(datalist) - 1  - i, v)
		}
	}

	{
		var datalist []TestObject
		for i:=0; i<6; i++ {
			o0 := &TestObjectInt{A:i}
			o1 := TestObject{A:i, B:float32(i), C:strconv.Itoa(i), D:o0}
			datalist = append(datalist, o1)
		}
		SortTestObject := func(left interface{}, right interface{}) bool {
			a := left.(TestObject)
			b := right.(TestObject)
			return a.A > b.A
		}

		SortInverse(datalist, SortTestObject)
		for i, v := range datalist {
			assert.Equal(t, i, v.A)
			assert.Equal(t, i, v.D.A)
		}

		Sort(datalist, SortTestObject)
		for i, v := range datalist {
			assert.Equal(t, len(datalist) - 1 - i, v.A)
			assert.Equal(t, len(datalist) - 1  - i, v.D.A)
		}
	}

	// 排序算法稳定性
	{
		var datalist []TestObject
		for i:=0; i<6; i++ {
			o0 := &TestObjectInt{A:i}
			o1 := TestObject{A:0, B:float32(i), C:strconv.Itoa(i), D:o0}
			datalist = append(datalist, o1)
			t.Logf("%d, %f", o1.A, o1.B)
		}
		var datalist2 []TestObject
		datalist2 = append(datalist2, datalist...)

		SortTestObject := func(left interface{}, right interface{}) bool {
			a := left.(TestObject)
			b := right.(TestObject)
			return a.A < b.A
		}
		Sort(datalist, SortTestObject)
		assert.Equal(t, datalist[1].B, float32(1))

		for i, v := range datalist {
			if i == len(datalist) - 1 {
				break
			}
			assert.Less(t, v.B, datalist[i+1].B)
		}

		SortInverse(datalist, SortTestObject)
		assert.Equal(t, datalist[len(datalist) - 1 - 1].B, float32(1))

		SortTestObject2 := func(left interface{}, right interface{}) bool {
			a := left.(TestObject)
			b := right.(TestObject)
			return a.A > b.A
		}
		Sort(datalist2, SortTestObject2)
		assert.Equal(t, datalist2[1].B, float32(1))
	}


}

