package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	iceCreams []string
}
type vehicule struct {
	doors int
	color string
}
type truck struct {
	vehicule  vehicule
	fourWheel bool
}
type sedan struct {
	vehicule vehicule
	luxury   bool
}

func main() {
	fmt.Println("EX1")
	p1 := person{
		firstname: "Sébastien",
		lastname:  "TACHIER",
		iceCreams: []string{"Chocolate", "Strawberry", "Vanilla"},
	}
	p2 := person{
		firstname: "Lucie",
		lastname:  "PICHON",
		iceCreams: []string{"Raspberry", "Gum", "Coconuts"},
	}

	fmt.Println("The favorite ice cream for " + p1.firstname + " " + p1.lastname + " are:")
	for i, v := range p1.iceCreams {
		fmt.Printf("\t%d - %v\n", (i + 1), v)
	}
	fmt.Println("The favorite ice cream for " + p2.firstname + " " + p2.lastname + " are:")
	for i, v := range p2.iceCreams {
		fmt.Printf("\t%d - %v\n", (i + 1), v)
	}
	fmt.Println("EX2")
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for _, v := range m {
		fmt.Println("The favorite ice cream for " + v.firstname + " " + v.lastname + " are:")
		for i, ic := range p2.iceCreams {
			fmt.Printf("\t%d - %v\n", (i + 1), ic)
		}
	}

	fmt.Println("EX3")
	t := truck{
		vehicule: vehicule{
			doors: 4,
			color: "Blue",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicule: vehicule{
			doors: 2,
			color: "Red",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.fourWheel)
	fmt.Println(s.luxury)
	fmt.Println("EX4")
	an := struct {
		format string
		grain  string
	}{
		format: "A4",
		grain:  "80",
	}
	fmt.Println(an.format)
	fmt.Println(an.grain)
}
