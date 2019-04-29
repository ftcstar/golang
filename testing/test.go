package main

import "fmt"

type intrf interface{}
type intrf1 interface{}

func checkfunc(a int, i intrf) {
	fmt.Println("Hello")
}

//var fnptr []func(int, interface{})
var fnptr []func(a int, i interface{})

func main() {
	fnptr = append(fnptr, checkfunc)
	fmt.Println(fnptr)
}
