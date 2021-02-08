package main

import "fmt"

// Bot est une interface qui déclare une
// function getGreeting qui retourne une chaine
type Bot interface {
	getGreeting() string
}

// EnglishBot is a bot
type EnglishBot struct{}

// SpanishBot is a bot
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	// On peut appeller printGreeting sur les deux objets car
	// Les deux implémentent l'interface Bot
	printGreeting(eb)
	printGreeting(sb)
}

// printGreeting atteint un objet de type Bot
func printGreeting(b Bot) {
	// Bot a une fonction getGreeting()
	fmt.Println(b.getGreeting())
}

// Note: si on utilise pas la variable du receiver, on peut
// L'ommettre à la déclaration et juste donner le type
func (EnglishBot) getGreeting() string {
	// Very Custom logic
	return "Hello there"
}

// Note: si on utilise pas la variable du receiver, on peut
// L'ommettre à la déclaration et juste donner le type
func (SpanishBot) getGreeting() string {
	// Very Custom logic
	return "Holla"
}
