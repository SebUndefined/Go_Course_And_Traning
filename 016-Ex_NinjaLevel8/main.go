package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}
type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

// SortPersonByAge sort
type SortPersonByAge []person

func (a SortPersonByAge) Len() int           { return len(a) }
func (a SortPersonByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortPersonByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func main() {
	fmt.Println("##################Ex1")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println(string(b))
	}
	fmt.Println("##################Ex2")
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var pers []person
	err2 := json.Unmarshal([]byte(s), &pers)
	if err2 != nil {
		fmt.Println("Erreur:", err2)
	} else {
		for i, v := range pers {
			fmt.Printf("The person in index %d in the slice is %v %v and he/she is %d years old and he/she says:\n", i, v.First, v.Last, v.Age)
			for _, v := range v.Sayings {
				fmt.Printf("\t- %v\n", v)
			}
		}
	}
	fmt.Println("##################Ex3")
	en := json.NewEncoder(os.Stdout)
	err3 := en.Encode(pers)
	if err3 != nil {
		fmt.Println("Someting is worg")
	}
	fmt.Println("##################Ex4")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println("##################Ex5")
	fmt.Println(pers)
	sort.Sort(SortPersonByAge(pers))
	fmt.Println(pers)
	sort.Slice(pers, func(i int, j int) bool {
		return pers[i].Last < pers[j].Last
	})
	fmt.Println(pers)
	sort.Slice(pers, func(i int, j int) bool {
		return pers[i].Last < pers[j].Last
	})
	for _, v := range pers {
		sort.Strings(v.Sayings)
	}
	fmt.Println(pers)
}
