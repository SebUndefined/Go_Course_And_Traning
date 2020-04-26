package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person person
	ltk    bool
}

// func (r receiver) identifier(parameters)(returns(s)){.....}
func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last)
}

func main() {
	foo()
	bar("seb")
	s1 := woo("Monneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
	variadic(1, 4, 5, 7, 8, 9, 2)

	xi := []int{1, 2, 3, 4, 5, 5, 6}
	sumXi := variadicReturn(xi...)
	fmt.Println("The total is", sumXi)
	defer deferFuncFoo()
	deferFuncBar()
	// Methods
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa2)
	sa2.speak()
	// Anonymous func
	func(x int) {
		fmt.Printf("Hello from anonymous func !!!. You have passed %d as argument\n", x)
	}(42)
	// Func expression
	f := func(x int) {
		fmt.Printf("My first func expression. %d\n", x)
	}
	f(44)
	// Returning func
	rs := returnString()
	fmt.Println(rs)
	fr := returnFunc()
	fmt.Printf("Type is %T\n", fr)
	frr := fr()
	fmt.Println("The result of the returned func is", frr)
	// Callback
	rc := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println("The sum is", rc)
	rc2 := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println("The sum of even is", rc2)
	rc3 := oddSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println("The sum of odd is", rc3)

	// Closure
	{
		var yc int
		yc = 1
		fmt.Println(yc)
	}
	// Do NOT WORK ==> fmt.Println(yc)
	icrem1 := incrementor()
	icrem2 := incrementor()
	fmt.Printf("Result if increm1 %d\n", icrem1())
	fmt.Printf("Result if increm1 %d\n", icrem1())
	fmt.Printf("Result if increm1 %d\n", icrem1())
	fmt.Printf("Result if increm1 %d\n", icrem1())
	fmt.Printf("Result if increm1 %d\n", icrem2())
	fmt.Printf("Result if increm1 %d\n", icrem2())
	fmt.Printf("Result if increm1 %d\n", icrem2())
	// Recursion
	rec := factorial(4)
	fmt.Println(rec)
}

// func (r receiver) identifier(parameters)(returns(s)){.....}
func foo() {
	fmt.Println("Hello from foo")
}

// Everything in go is pass by value !!!!
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
}

func variadic(x ...int) {
	fmt.Println("Variadic params")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		fmt.Println("For item in position", i, "we are adding", v, "to the total", sum)
		sum += v
	}
	fmt.Println("The total is", sum)
}
func variadicReturn(x ...int) int {
	fmt.Println("Variadic params")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		fmt.Println("For item in position", i, "we are adding", v, "to the total", sum)
		sum += v
	}
	return sum
}

func deferFuncFoo() {
	fmt.Println("Defer func foo")
}
func deferFuncBar() {
	fmt.Println("Defer func bar")
}

func returnString() string {
	s := "Hello world"
	return s
}
func returnFunc() func() int {
	return func() int {
		return 451
	}
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func evenSum(f func(xi ...int) int, n ...int) int {
	var numbers []int
	for _, v := range n {
		if v%2 == 0 {
			numbers = append(numbers, v)
		}
	}
	return f(numbers...)
}

func oddSum(f func(xi ...int) int, n ...int) int {
	var numbers []int
	for _, v := range n {
		if v%2 != 0 {
			numbers = append(numbers, v)
		}
	}
	return f(numbers...)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n == 0 {
		fmt.Println(n)
		return 1
	}
	fmt.Println(n)
	return n * factorial(n-1)
}
