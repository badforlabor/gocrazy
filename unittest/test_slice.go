/**
 * Auth :   liubo
 * Date :   2019/12/7 17:37
 * Comment:
 */

package main

import (
	"github.com/stretchr/testify/assert"
	. "github.com/badforlabor/gocrazy/alg"
	"strconv"
)

func TestSlice1(t assert.TestingT) {
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

	assert.Equal(t, 0, 1, "")
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

func TestSlice2(t assert.TestingT) {
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
}

func TestClone(t assert.TestingT) {
	{
		var datalist = []int{0,1,2,3,4,5}
		var clone1 []int
		Clone(&clone1, datalist)
		assert.Equal(t, len(clone1),len(datalist))
		for i:=0; i<len(datalist); i++ {
			assert.Equal(t, datalist[i], clone1[i])
		}

		SliceInverse(&clone1)
		for i:=0; i<len(datalist); i++ {
			assert.Equal(t, datalist[i], clone1[len(datalist)-1-i])
		}
	}
}
