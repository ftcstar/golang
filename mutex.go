package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mx sync.Mutex
var sumV = 0

func sum(name string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//mx.Lock()
			sumV = sumV + 1
			//mx.Unlock()

		}
		fmt.Println("From "+name+":", sumV)
	}()
}

func main() {
	p := []string{"p1", "p2", "p3"}
	for _, name := range p {
		fmt.Println("Main>> ", name)
		sum(name)
	}
	wg.Wait()
	fmt.Println("Final sum:", sumV)
}
