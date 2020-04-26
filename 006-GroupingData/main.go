package main

import "fmt"

func main() {
	// ARRAY
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	// SLICE
	w := []int{3, 5, 8, 9, 12}
	fmt.Println(w)
	// Update index
	w[0] = 200
	// LOOP
	for i, v := range w {
		fmt.Printf("Index %d is %d \n", i, v)
	}
	// Slicing a slice
	// start at an index until the end
	fmt.Println(w[1:])
	// start at an index until another index BUT NOT INCLUDING THE INDEX !!!
	fmt.Println(w[1:3])

	// APPEND A SLICE
	w = append(w, 22, 10, 1988)
	fmt.Println(w)
	// variadic
	y := []int{}
	y = append(y, w...)
	fmt.Println(y)
	// DELETE FROM SLICE
	w = append(w[:2], w[4:]...)
	fmt.Println(w)
	// SLICE MAKE
	m := make([]int, 10, 12)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m[0] = 42
	m[3] = 45
	//Do not work m[10] = 13
	m = append(m, 9876)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m = append(m, 9877)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m = append(m, 666)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	// Multidimentionnal
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	// MAP
	u := map[string]int{
		"James":       32,
		"Monneypenny": 27,
	}
	fmt.Println(u)
	fmt.Println(u["James"])
	// Not in the map, return 0
	fmt.Println(u["Seb"])
	// comma ok idiom
	v, ok := u["Seb"]
	if !ok {
		fmt.Println("No entry in the map with key Seb")
	} else {
		fmt.Println("Entry found")
	}
	fmt.Println(ok)
	fmt.Println(v)
	// ADD an element
	u["todd"] = 34
	fmt.Println(u)
	// LOOP
	for k, v := range u {
		fmt.Println(k, v)
	}
	// DELETE
	delete(u, "todd")
	fmt.Println(u)
	// Delete unexisting key...
	delete(u, "Seb")
	fmt.Println(u)
	// comma ok idiom
	if v, ok := u["Monneypenny"]; ok {
		fmt.Println("Deleting key ", "Monneypenny", v)
		delete(u, "Monneypenny")
	}
	fmt.Println(u)
}
