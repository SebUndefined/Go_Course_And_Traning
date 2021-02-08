package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Qui est un slide de string
type deck []string

// Fonction pour créer un nouveau deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Opt 1 log error et retuirn a call to newDeck()
		// Opt 2 Log the error et quitter le programme
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Convertion en slice de string (après avoir convertie en chaine de caractère)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// d est un receiver
// Toutes les variables de type deck ont accès
// à la méthode print
// d est une copie de deck (dans ce cas) avec lequel on a appellé print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Pas de receiver car ambigue
// Deal, split a deck depending of the handsize
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	// Create a source. La source dépendra du Unix time
	source := rand.NewSource(time.Now().UnixNano())
	// Create a rand
	r := rand.New(source)

	for i := range d {
		// Trouver un nombre entre 0 et la longueur du slice -1
		newPosition := r.Intn(len(d) - 1)
		// Take whatever is at d[newPosition] and assign it to d[i]
		// And
		// Take whatever is at d[i] and assign it to d[newPosition]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	// Convertion de d en []string puis join avec le package strings
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
