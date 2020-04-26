package main

import "fmt"

func main() {
	fmt.Println("EX1")
	for i := 1; i <= 10000; i++ {
		fmt.Printf("Current number is %d \n", i)
	}
	fmt.Println("EX2")
	for i1 := 65; i1 <= 90; i1++ {
		fmt.Printf("Current number is %d \n", i1)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t %#U \n", i1)
		}
	}
	fmt.Println("EX3")

	fmt.Println("EX4")
	fmt.Println("EX5")
	for i := 10; i <= 100; i++ {
		fmt.Println((i % 4))
	}
	fmt.Println("EX6 EX7")
	x := 54
	if x == 43 {
		fmt.Println("Hey, it's 43 !!")
	} else if x <= 53 {
		fmt.Println("I am sure it's less or equal to 53 ?!")
	} else {
		fmt.Println("Ok can not guess it...")
	}
	fmt.Println("EX8")
	name := "SebUndefined"
	switch name {
	case "Bond":
		fmt.Println("OJames ?! Is that you ?!")
	case "SebUndefined":
		fmt.Println("You are Sebundefined")
	default:
		fmt.Println("Don't know you !")
	}

	fmt.Println("EX9")
}
