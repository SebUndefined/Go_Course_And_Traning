package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	fname := flag.String("filename", "", "The filename you want to read")
	flag.Parse()
	if len(*fname) <= 0 {
		fmt.Println("The string length has to be upper than 0.")
		os.Exit(1)
	}
	fmt.Println(*fname)
	f, err := os.Open(*fname)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
