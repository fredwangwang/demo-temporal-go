package main

import (
	"fmt"
	"reflect"
)

type IFunc interface {
	Func(b string)
}

type A struct {
}

func (a *A) Func(b string) {
	fmt.Println(b)
}

func main() {
	a := A{}

	funcType := reflect.TypeOf(a.Func)
	ifaceType := reflect.TypeOf(IFunc.Func)

	ifaceTypeNumin := ifaceType.NumIn()

	fmt.Printf("%d, %d\n", funcType.NumIn(), ifaceTypeNumin) // 1, 2

	x := IFunc.Func
	x(&a, "ab")

	ifaceVal := reflect.ValueOf(IFunc.Func)
	ifaceVal.Call([]reflect.Value{reflect.ValueOf(&a), reflect.ValueOf("aaa")})
}
