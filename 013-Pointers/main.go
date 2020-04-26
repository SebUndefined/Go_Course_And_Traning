package main

import "fmt"

type person struct {
	last  string
	first string
	age   int
}

func main() {
	a := 42
	fmt.Println("Show var a and (&) the pointer to the adress")
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("The type of a %T\n", a)
	fmt.Printf("The type of &a %T\n", &a)

	var b *int = &a
	fmt.Println(b)
	// Show value store at the adress
	fmt.Println(*b)
	*b = 54
	fmt.Println(a)
	fmt.Println(*b)

	x := 0
	foo(&x)
	fmt.Println(x)
	fmt.Println("RECEIVERS")
	p := person{
		first: "James",
		last:  "Bond",
		age:   31,
	}
	fmt.Println(p)
	p.setLast("Bourne")
	p.setFirst("Jason")
	p.setAge(35)
	fmt.Println(p)
}

func foo(x *int) {
	fmt.Println(x)
	*x = 43
	fmt.Println(x)
	fmt.Println(*x)
}

func (p *person) setLast(s string) {
	p.last = s
}
func (p *person) setFirst(s string) {
	p.first = s
}
func (p *person) setAge(a int) {
	p.age = a
}
