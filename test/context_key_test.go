package tests

import (
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
	"reflect"
)

var srv *tester.Tester

const (
	normalUserIdKey       = "user/normal/key"
	normalUserPassword    = "yY11112222"
	normalUserNewPassword = "xX11122222"
)

var intT = 1
var intType = reflect.TypeOf(intT)
var uintType = reflect.TypeOf(uint(1))

func RangeInt(l, r int) []int {
	var x = make([]int, r-l)
	for i := l; i < r; i++ {
		x[i-l] = i
	}
	return x
}
