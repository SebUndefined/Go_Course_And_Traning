package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("[INFO] Starting program")
	if len(os.Args) <= 1 {
		fmt.Println("No file profided")
		os.Exit(1)
	}
	fn := os.Args[1]
	fmt.Println("[INFO] Looking for file " + fn)
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("[ERROR] ", err)
		os.Exit(1)
	}
	fmt.Println("[INFO] Writing file content")
	io.Copy(os.Stdout, f)
	fmt.Println("[INFO] End")
}
