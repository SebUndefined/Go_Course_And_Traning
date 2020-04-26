package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person person
	ltk    bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}
	sa2 := secretAgent{
		person: person{
			first: "MM",
			last:  "Steps",
			age:   23,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(sa1.person.first, sa1.person.last)
	fmt.Println(sa2.person.first, sa2.person.last)

}
