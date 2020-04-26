package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC FOR GENERATING GREETING
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	// VERY CUSTOM LOGIC FOR GENERATING GREETING
	return "Hola senor"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
