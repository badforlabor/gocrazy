/**
 * Auth :   liubo
 * Date :   2019/12/6 17:56
 * Comment:
 */

package alg

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSlice1(t *testing.T) {
	var datalist = []int{0,1,2,3,4,5}

	SliceRemoveAt(&datalist, 1)
	assert.Equal(t, len(datalist), 5, "")

	SliceRemoveOne(&datalist, IntCompare(0))
	assert.Equal(t, datalist[0], 2, "")

	SliceAddOne(&datalist, 6)

	SliceInsertOne(&datalist, 0, 0)
	SliceInsertOne(&datalist, 1, 1)

	SliceAddUnique(&datalist, 6, IntCompare(6))
	SliceAddUnique(&datalist, 7, IntCompare(7))

	assert.Equal(t, len(datalist), 8, "")

	for i:=len(datalist); i<32; i++ {
		SliceAddUnique(&datalist, i, IntCompare(i))
	}

	for i, v := range datalist {
		assert.Equal(t, v, i, "")
	}
}

func TestSlice2(t *testing.T) {
	var datalist []TestObject
	for i:=0; i<6; i++ {
		o0 := &TestObjectInt{A:i}
		o1 := TestObject{A:i, B:float32(i), C:strconv.Itoa(i), D:o0}
		datalist = append(datalist, o1)
	}

	for i:=len(datalist); i<32; i++ {
		o0 := &TestObjectInt{A:i}
		o1 := TestObject{A:i, B:float32(i), C:strconv.Itoa(i), D:o0}
		datalist = append(datalist, o1)
		SliceAddUnique(&datalist, i, o1.Compare())
	}

	for i, v := range datalist {
		if v.A == i && v.C == strconv.Itoa(i) && v.D.A == i {

		} else {
			assert.Fail(t, "failed.", "failed:%d", i)
		}
	}

	findid := 10
	findfunc := func(d interface{}) bool {
		v := d.(TestObject)
		return v.A == findid
	}
	idx := Find(datalist, findfunc)
	assert.Equal(t, idx, findid)

	SliceRemoveOne(&datalist, findfunc)
	idx = Find(datalist, findfunc)
	assert.Equal(t, idx, -1)
}

func TestSlice3(t *testing.T) {
	{
		var datalist = []int{0,1,2,3,4,5}
		SliceInverse(datalist)
		for i:=0; i<len(datalist); i++ {
			assert.Equal(t, i, datalist[len(datalist)-1-i])
		}
	}
	{
		var datalist = []int{0,1,2,3,4}
		SliceInverse(datalist)
		for i:=0; i<len(datalist); i++ {
			assert.Equal(t, i, datalist[len(datalist)-1-i])
		}
	}
}

type TestObjectInt struct {
	A int
}
type TestObject struct {
	A int
	B float32
	C string
	D *TestObjectInt
}
func (self *TestObject) Compare() func(d interface{})bool {
	var function = func(d interface{})bool {
		var aa = d.(TestObject)
		return *self == aa
	}
	return function
}
