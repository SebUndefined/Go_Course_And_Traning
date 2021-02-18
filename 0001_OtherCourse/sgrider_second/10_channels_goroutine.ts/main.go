package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Déclaration des liens à surveiller
	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com",
	}
	// Créer un channel de type string
	c := make(chan string)
	// Loop sur la liste pour checker les sites les uns après les autres
	for _, link := range links {
		// On check le lien et on passe le channel en paramètre
		go checkLink(link, c)
	}
	// Equivalent à ce qu'il y a au dessus.
	// On dit "attend le channel et dès que tu as une valeur
	// Lance une nouvelle goroutine de la fonction litteral
	for l := range c {
		// Invoquer que fonction litteral
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// On ne se souci pas de resp ("_") et on prend simplement l'erreur
	_, err := http.Get(link)
	// S'il y a une erreur, on affiche que le site est probablement down
	if err != nil {
		fmt.Println(link, " might be down")
		// Renvoyer le lien au channel
		c <- link
		return
	}
	// Sinon, on affiche que le site est up
	fmt.Println(link, " is up")
	// Renvoyer le lien au channel
	c <- link
}
