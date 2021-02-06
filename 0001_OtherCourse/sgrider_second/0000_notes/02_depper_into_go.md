# Deeper into go

## Déclaration de variable

```go
package main

import "fmt"

func main() {
    // Simple (nom + type)
    var card string
    // Simple (nom + type + assignation)
    var card string = "Ace of spades"
    // Inférence de type, Go détermine le type avec l'asignation (nom + assignation)
    var card = "Ace of spades"
    // Inférence mais courte. Utilisé uniquement lors de déclaration de nouvelles variables. 
    card := "Ace of spades"
	fmt.Println(card)
}

```



## Les types en go

### Types basics

- **bool:** true ou false
- **string:** Chaine de caractère
- **int:** un entier
- **float64:** nombre à virgule

## Déclaration de function

```go
package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// func nomFonction(param1 type, param2 type...) typeDeRetour
func newCard() string {
	return "Five of Diamonds"
}

```

## Slices et boucles for

### Arrays et slices en go

|        | Array                                       | Slices                                      |
| ------ | ------------------------------------------- | ------------------------------------------- |
| Taille | Taille fixe d'une liste de "chose"          | Tableau qui peut grandir et rétrécir        |
| Typage | Tous les éléments doivent être du même type | Tous les éléments doivent être du même type |
|        |                                             |                                             |
|        |                                             |                                             |

```go
package main

import "fmt"

func main() {
	// Déclarer le slice
	cards := []string{"Ace of Diamonds", newCard()}
	// Ajouter un élémenent au slice. Cela ne modifie pas le slice
	// existant mais retourne un nouveau
	// slice que l'on peut assigner à la variable souhaité.
	// (Spécificité de go)
	cards = append(cards, "Six of Spades")
	// i = index courant
	// card = card courante
	// cards = le slice à parcourir
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
```

## Approche orienté objet vs approche de go

Go n'est pas un language de programmation orienté objet. Il a sa propre façon d'aborder le problème. 

L'idée est de prendre un type "basique" (string, integer, float, map...) et d'étendre ses fonctionnalités. 

deck.go
```go
package main

import "fmt"

// Create a new type of 'deck'
// Qui est un slide de string
type deck []string

// d est un receiver
// Toutes les variables de type deck ont accès
// à la méthode print
// d est une copie de deck (dans ce cas) avec lequel on a appellé print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

```

main.go
```go
package main

func main() {
	// Déclarer le slice
	cards := deck{"Ace of Diamonds", newCard()}
	// Ajouter un élémenent au slice. Cela ne modifie pas le slice
	// existant mais retourne un nouveau
	// slice que l'on peut assigner à la variable souhaité.
	// (Spécificité de go)
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

```

### Créer un nouveau deck (constructeur en quelque sorte)

deck.go
```go
package main

import "fmt"

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
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
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

```


main.go
```go
package main

func main() {
	// Déclarer le slice
	cards := newDeck()

	cards.print()
}
```

## Slice range syntax

Les slices sont zero based syntax.
```go
fruit := []string{"apple", "banana", "grappe", "orange"}

// fruit[startIndexIncluding:upToNotIncluding]
fruit[0:2] // apple et banana
fruit[:2] // apple et banana
fruit[2:] // grappe et orange
```

Exemple avec le deck. Deal function
```go
// Deal, split a deck depending of the handsize
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Utilisation dans main.go
hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
```

## Byte Slices

Byte slice représente une chaine de caractère. Il représente une chaine sous forme décimal (slice de chiffre)

## Type conversion

Il s'agit d'une grande spécificité de go. La convertion est facile et il est possible d'implementer ses propres convertions. Exremple string vers byte slice

```go
myString := "Hi there"
fmt.Println([]byte(myString))
// Output [72 105 32 116 104 101 114 101]
```


## Test avec go

Go testing n'utilise pas de framework. On dispose en revanche d'une petite lib native. 

Il suffit de créer un fichier se terminant par _test.go

Pour lancer les tests go test