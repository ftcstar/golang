package main

import (
	"fmt"
	"time"
)

func testR() {
	defer recvr()
	fmt.Println("testR is in progress")
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		panic("hello")
	}
	panic("testR completed")
}

func testR1() {
	fmt.Println("testR1 is in progress")
	for i := 0; i < 10; i++ {
		time.Sleep(5)
	}
	fmt.Println("testR1 completed")
}

func recvr() {
	r := recover()
	fmt.Println("-----", r)
}

func main() {
	go testR()
	go testR1()
	fmt.Println("In Main")
	for {
	}
}
