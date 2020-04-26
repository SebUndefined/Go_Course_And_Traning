package main

import "fmt"

func main() {
	// LOOP
	// Simple loop
	for i := 0; i <= 10; i++ {
		for y := 0; y <= 5; y++ {
			fmt.Printf("The inner loop %d\t The outer loop %d\n", y, i)
		}
		fmt.Println("")
	}

	// Break keyword
	x := 0
	for {
		if x > 4 {
			break
		}
		fmt.Println(x)
		x++
	}

	// Break continue keyword
	fmt.Println("Break continue keyword")
	x = 0
	for {

		if x > 100 {
			break
		} else {
			if x%2 != 0 {
				x++
				continue
			} else {
				fmt.Println(x)
				x++
			}
		}
	}
	// Print ASCII
	for counter := 33; counter <= 56; counter++ {
		fmt.Printf("%v\t %#U \n ", counter, counter)
	}

	// IF
	x = 42
	if x == 42 {
		fmt.Println("Our value is 42")
	}

	if x <= 40 {
		fmt.Println("value less that 40")
	} else {
		fmt.Println("value more that 40")
	}

	if x <= 42 && x%2 == 0 {
		fmt.Println("OK")
	}

	// Switch
	switch {
	case false:
		fmt.Println("this souhld not print")
	case (2 == 4):
		fmt.Println("this souhld not print")
	case (3 == 3):
		fmt.Println("this souhld print")
	case (4 == 4):
		fmt.Println("this souhld not print because above condition is statisfied")
	}
	// Switch fallthrough... Do not use it.
	switch {
	case false:
		fmt.Println("this souhld not print")
	case (2 == 4):
		fmt.Println("this souhld not print")
	case (3 == 3):
		fmt.Println("this souhld print")
		fallthrough
	case (4 == 5):
		fmt.Println("this souhld print because we use fallthrough")
		fallthrough
	case (4 == 6):
		fmt.Println("this souhld print because we use fallthrough")
	default:
		fmt.Println("default")

	}
}
