package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters)(returns(s)){.....}
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "the secret agent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I am a human", h)
	switch h.(type) {
	case secretAgent:
		fmt.Println("and a secret agent")
	case person:
		fmt.Println("and a personn")
	default:
		fmt.Println("and undefined")
	}
}

func main() {
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
	p1 := person{
		first: "Dr.",
		last:  "No",
	}
	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)
	p1.speak()
	// Conversion

}
