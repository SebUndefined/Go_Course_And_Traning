package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) changeMe() {
	p.first = "Jason"
	p.last = "Bourne"
	p.age = 31
}
func changeMe(p *person) {
	p.first = "Nicky"
	p.last = "Larson"
	p.age = 27
}

func main() {
	x := "This is a string"
	fmt.Println(&x)

	p := person{
		first: "James",
		last:  "Bourne",
		age:   45,
	}
	fmt.Println(p)
	p.changeMe()
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
