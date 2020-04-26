package main

import "fmt"

func main() {
	// ARRAY
	fmt.Println("Ex1")
	a := [5]int{22, 22, 34, 54, 32}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("The type of the var a is %T", a)
	fmt.Println("Ex2")
	s := []int{1, 12, 22, 43, 56, 22, 890, 1, 4, 32}
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Printf("The type of the var s is %T", s)
	fmt.Println("Ex3")
	se := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(append(se[:5]))
	fmt.Println(append(se[5:]))
	fmt.Println(append(se[2:7]))
	fmt.Println(append(se[1:6]))
	fmt.Println("Ex4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("EX5")

	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = []int{}
	y = append(x[:3], x[6:]...)
	fmt.Println(y)
	fmt.Println("Ex6")
	ss := []string{`Alabama`, `Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
		` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`,
		` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(ss))
	fmt.Println(cap(ss))
	for i := 0; i < len(ss); i++ {
		fmt.Printf("State %v is in position %d \n", ss[i], i)
	}
	fmt.Println("Ex7")
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	famous := [][]string{jb, mp}
	for _, t := range famous {
		for _, p := range t {
			fmt.Println(p)
		}
	}
	fmt.Println("Ex8")
	fmaousMap := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range fmaousMap {
		fmt.Printf("Key %v \n", k)
		for i, f := range v {
			fmt.Printf("\tIndex %v has value %v\n", i, f)
		}
	}
	fmt.Println("Ex9")
	fmaousMap["seb_me"] = []string{"Lazy", "jj", "Beer", "Sport"}
	for k, v := range fmaousMap {
		fmt.Printf("Key %v \n", k)
		for i, f := range v {
			fmt.Printf("\tIndex %v has value %v\n", i, f)
		}
	}
	fmt.Println("Ex10")
	_, ok := fmaousMap["seb_me"]
	if ok {
		delete(fmaousMap, "seb_me")
		fmt.Println("Value deleted")
	}
	for k, v := range fmaousMap {
		fmt.Printf("Key %v \n", k)
		for i, f := range v {
			fmt.Printf("\tIndex %v has value %v\n", i, f)
		}
	}
}
