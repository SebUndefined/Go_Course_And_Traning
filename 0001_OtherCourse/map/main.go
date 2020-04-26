package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//var colors map[string]string
	//colors := make(map[int]string)

	// Add value
	//colors[101] = "#ffffff"
	// delete value
	//delete(colors, 101)

	// Iterate over key value

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hexCode := range c {
		fmt.Println("Hex code for " + color + " is " + hexCode)
	}
}
