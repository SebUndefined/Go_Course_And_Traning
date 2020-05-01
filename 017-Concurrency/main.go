package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("OS %v\n", runtime.GOOS)
	fmt.Printf("ARCH %v\n", runtime.GOARCH)
	fmt.Printf("CPUs %v\n", runtime.NumCPU())
	fmt.Printf("Goroutine %v\n", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Printf("OS %v\n", runtime.GOOS)
	fmt.Printf("ARCH %v\n", runtime.GOARCH)
	fmt.Printf("CPUs %v\n", runtime.NumCPU())
	fmt.Printf("Goroutine %v\n", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
