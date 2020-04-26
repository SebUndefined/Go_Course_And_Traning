package main

import "fmt"

const ca = 42
const cb string = "Pouet"
const (
	i1 = iota + 2021
	i2
	i3
	i4
)

func main() {
	y := 255
	fmt.Printf("%d, %b, %#x", y, y, y)
	fmt.Println("EX 2")
	eq := (y == 255)
	lesseq := (y <= 234)
	upeq := (y >= 255)
	dif := (y != 234)
	inf := (y < 300)
	sup := (y > 234)

	fmt.Println("", eq, lesseq, upeq, dif, inf, sup)
	fmt.Println("EX 3")
	fmt.Println(ca)
	fmt.Println(cb)
	fmt.Println("EX 4")
	bsTest := 56
	fmt.Printf("%d, %b, %#x \n", bsTest, bsTest, bsTest)
	bsTest2 := bsTest << 1
	fmt.Printf("%d, %b, %#x", bsTest2, bsTest2, bsTest2)
	fmt.Println("EX 5")
	s := `This is
	a 
	raw string 
	literal`
	fmt.Println(s)
	fmt.Println("EX 6")
	fmt.Printf("%d, %d, %d, %d", i1, i2, i3, i4)
}
