package main

// https://mholt.github.io/json-to-go/
// $ go get golang.org/x/crypto/bcrypt ==> get the package
// $ go get -u golang.org/x/crypto/bcrypt ==> update the package
import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

//
// Person structure
type Person struct {
	Last  string
	First string
	Age   int
}

// ByAge sort a slice of person by age
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := Person{
		First: "Miss",
		Last:  "Monneypenny",
		Age:   23,
	}
	p3 := Person{
		First: "Q",
		Last:  "S",
		Age:   32,
	}
	p4 := Person{
		First: "Dr",
		Last:  "No",
		Age:   99,
	}
	people := []Person{p1, p2, p4, p3}
	fmt.Println(people)
	sheet, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(sheet))
	}
	var people2 []Person
	json.Unmarshal(sheet, &people2)
	fmt.Println(people2)
	// SOrt
	xi := []int{55, 32, 12, 8, 104, 45, 31}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	xs := []string{"James", "Seb", "Anabelle", "Yann", "Lucie"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Last < people[j].Last
	})
	fmt.Println(people)

	// Bcrypt
	pass := "Amazing Password"
	pb, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	fmt.Println(pass)
	fmt.Println(pb)
	errResult := bcrypt.CompareHashAndPassword(pb, []byte(pass))
	if errResult != nil {
		fmt.Println("Erreur:", errResult)
	} else {
		fmt.Println("all good")
	}

	errResult2 := bcrypt.CompareHashAndPassword(pb, []byte("Password Fake"))
	if errResult2 != nil {
		fmt.Println("Erreur:", errResult2)
	} else {
		fmt.Println("all good")
	}

}
