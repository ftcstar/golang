package main

import "fmt"

func main() {
	a := [...]int{123, 20, 30, 40}
	a[3] = 20
	fmt.Println(a)

	for k, v := range a {
		fmt.Printf("a[%d]=%d\n", k, v)
	}
}
