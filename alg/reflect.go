/**
 * Auth :   liubo
 * Date :   2019/12/3 16:10
 * Comment: 各种反射的应用
 */

package alg

import "reflect"

func mustPtr(dataList interface{}) {
	s := reflect.ValueOf(dataList)

	if s.Kind() != reflect.Ptr {
		panic("must ptr. please use operator &")
	}
}