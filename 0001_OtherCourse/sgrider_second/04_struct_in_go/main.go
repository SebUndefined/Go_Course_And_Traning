package main

import "fmt"

// ContactInfo is
type ContactInfo struct {
	email   string
	zipCode int
}

// Person struct
type Person struct {
	firstname string
	lastname  string
	// Embedded struct
	contact ContactInfo
}

func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
	}}

	// Go est assez malin pour prendre le pointer si le receiver le demande
	// Il prendra le pointer automatiquement
	jim.updateName(("Jimmy"))
	jim.print()
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
