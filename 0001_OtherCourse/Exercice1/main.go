package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceOfInts := getSliceOfInts(10)

	for _, n := range sliceOfInts {
		if n%2 == 0 {
			fmt.Println(strconv.Itoa(n) + " is even")
		} else {
			fmt.Println(strconv.Itoa(n) + " is odd")
		}
	}
}

func getSliceOfInts(max int) []int {
	var sliceOfInts []int
	for i := 0; i <= max; i++ {
		sliceOfInts = append(sliceOfInts, i)
	}
	return sliceOfInts
}
