package main

import "fmt"

var z = 43
var a string

type hotdog int

var b hotdog

func main() {
	n, err := fmt.Println("Hello world", "Pouet", 45, true)
	fmt.Println(n)
	fmt.Println(err)
	// Short declaration operator
	x := 42
	fmt.Println(x)
	// Assign 99 to x
	x = 99
	fmt.Println(x)
	// Long declaration
	var y string = "test"
	fmt.Println(y)
	foo()
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	x = 156
	fmt.Printf("%#x\n", x)

	b = 44
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	z = int(b)
	fmt.Printf("%v\n", z)
	fmt.Printf("%T\n", z)
}

func foo() {
	fmt.Println(z)
	fmt.Println(a)
	a = "Héhé"
	fmt.Println(a)
}
