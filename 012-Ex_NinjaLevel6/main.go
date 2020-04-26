package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last  string
	age   int
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func main() {
	fmt.Println("################ Ex1")
	fmt.Println(fooEx1())
	xEx1, yEx1 := barEx1()
	fmt.Println(xEx1, yEx1)
	fmt.Println("################ Ex2")
	nEx2 := []int{10, 20, 30, 40}
	fmt.Println(fooEx2(nEx2...))
	fmt.Println(barEx2(nEx2))
	fmt.Println("################ Ex3")
	defer deferFuncEx3()
	fmt.Println("################ Ex4")
	pEx4 := person{
		first: "Sébastien",
		last:  "Tachier",
		age:   31,
	}
	pEx4.speak()
	fmt.Println("################ Ex5")
	sq := square{
		length: 4.34,
	}
	ci := circle{
		radius: 3.45,
	}
	fmt.Printf("The area of the square is %f\n", sq.area())
	fmt.Printf("The area of the circle is %f\n", ci.area())
	info(sq)
	info(ci)
	fmt.Println("################ Ex6")
	func(i int) {
		fmt.Println("This anonmymous function print the int value passed to it:", i)
	}(31)
	fmt.Println("################ Ex7")
	fuEx7 := func(i int) {
		fmt.Println("This anonmymous function print the int value passed to it:", i)
	}
	fuEx7(44)
	fmt.Println("################ Ex8")
	fuEx8 := iReturnFunc()
	fuEx8(666)
	fmt.Println("################ Ex9")
	iGetCallback(fuEx8)
}

func fooEx1() int {
	return 51
}
func barEx1() (int, string) {
	return 51, "Hello World!!"
}
func fooEx2(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
func barEx2(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
func deferFuncEx3() {
	fmt.Println("Defer func Ex3 should be executed after all other ;)")
}
func (p person) speak() {
	fmt.Printf("Hello my name is %v %v and I am %d years old\n", p.first, p.last, p.age)
}
func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * 2
}
func info(s shape) {
	fmt.Printf("%f\n", s.area())
}
func iReturnFunc() func(i2 int) {
	return func(i2 int) {
		fmt.Println("This function is returned by the iReturnFunc and return an int which is", i2)
	}
}
func iGetCallback(f func(i int)) {
	f(44000)
}
