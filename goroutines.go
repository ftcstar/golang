package main

import (
	"fmt"
	"sync"
)

/*
var wg sync.WaitGroup

func r1() {
	defer wg.Done()
	fmt.Println("r1>> started")
	fmt.Println("r1>> completed")
}

func r2() {
	defer wg.Done()
	fmt.Println("r2>> started")
	fmt.Println("r2>> completed")
}

func main() {
	fmt.Println("Started main")
	wg.Add(2)
	go r1()
	go r2()
	wg.Wait()
}
*/

var wg sync.WaitGroup

func r1(c chan int) {
	defer wg.Done()
	fmt.Println("R1 started")
	c <- 10
	fmt.Println("R1 completed")
}

func r2(c chan int) {
	defer wg.Done()
	value := <-c
	if value == 10 {
		fmt.Println("Received value:", value)
	}
	fmt.Println("R2 completed")
}

func main() {
	c := make(chan int)
	wg.Add(2)
	go r2(c)
	go r1(c)
	wg.Wait()
}
