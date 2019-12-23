/**
 * Auth :   liubo
 * Date :   2019/12/7 20:22
 * Comment:
 */

package alg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
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
