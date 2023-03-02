package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	var x = 1024
	var str = "HaiCoder"
	pValueX := reflect.ValueOf(&x)
	pValueStr := reflect.ValueOf(&str)
	fmt.Println("pvalueX =", pValueX)
	fmt.Println("pvalueStr =", pValueStr)
	pValueElemX := pValueX.Elem()
	pValueElemStr := pValueStr.Elem()
	fmt.Println("pValueElemX =", pValueElemX)
	fmt.Println("pValueElemStr =", pValueElemStr)
}
