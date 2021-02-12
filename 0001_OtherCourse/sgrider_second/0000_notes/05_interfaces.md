# Interfaces

## Utilité des interfaces

On sait que:
- Toutes les valeurs ont un type
- Toutes les fonctions doivent spécifier le type de leurs arguments

**Donc, cela siginifie t'il que toutes les fonctions que nous écrivons doivent être écrite pour convenir à différents types même si la logique est la même ?**

### Problème sans les interfaces

Comme précisé ci-dessous la surcharge de méthode est interdite en go. De plus, les méthodes printGreeting sont très similaire en logique. Sans interfaces, il faudrait déclarer deux méthodes, une pour chaque struct

```go
package main

import "fmt"

// EnglishBot is a bot
type EnglishBot struct{}

// SpanishBot is a bot
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb EnglishBot) {
	fmt.Println(eb.getGreeting())
}
// Non permis, surcharge de méthode non autorisée en go !!!
func printGreeting(sb SpanishBot) {
	fmt.Println(sb.getGreeting())
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

```

### Interfaces en pratique

```go
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
```

Note à propos des interfaces

- **Interfaces ne sont pas des types génériques**: Les autres languages ont des types génériques, go non
- **Les interfaces sont implicites**: On ne précise pas manuellement que notre type custom satisfait l'interface. 
- **Les interfaces sont des contrats et nous aide à gérer nos types**: Si l'implémentation d'une fonction de notre type custom est cassée, les interfaces ne nous aideront pas. 
- **Les interfaces sont dures, step 1: les comprendre**: Comprendre les interfaces en comment les lire dans une librairie standard. Ecrire ses propres interfaces est dur et demande de l'expérience. 



## Illustration avec HTTP

Utilisation de la lib standard HTTP

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
```

Dans ce cas de figure, on peut remarquer que l'on affiche pas le body de la requête. On peut néanmoins remarquer que le type de retour est un pointer vers un objet de type io.ReadCloser... 

```go
Body io.ReadCloser
```

Si on analyse ce type sur la doc go, on peut voir que ce type à un field Body de type io.ReadCloser. La syntaxe vient du fait que dans ce cas, on "assemble" les interfaces.

```go
type ReadCloser interface {
    Reader
    Closer
}
```

ReadCloser est une interface qui a deux fields Reader et Closer. Reader est une simple interface qui a une fonction Read qui prend un tableau de byte en paramètre. 

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Closer a une simple methode Close()

```go
type Closer interface {
    Close() error
}
```

Pourquoi ? 

Cela permet de généraliser !! Ce que cela veut dire: quel que soit le type de body, tant qu'il satisfait l'interface ReadCloser. Met ce que tu veux dans body, tant qu'il satisfait la ReadCloser Interface.

**Imaginons cela sans interface**

| Source of input                              | Return    | To Print it               |
| -------------------------------------------- | --------- | ------------------------- |
| HTTP Request Body                            | []flargen | func printHTTP([]flargen) |
| Text file on HD                              | []string  | func printFile([]string)  |
| Image From HD                                | jpegne    | func printImage(jpegne)   |
| User input from CLI                          | []byte    | func printText([]byte)    |
| Data from analog sensor plugged into machine | []float   | func printData([]float)   |

Avec tous ces types, on devrait définir une fonction différente pour chacun de ces types. 

L'idée de la fonction ```Read``` de l'interface Reader est passe moi un []byte et je vais le remplir avec les données nécessaire (pointer, modifie la valeur). Je te retourne un int (n de byte écrit) et une err.

Exemple avec HTTP

```go

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.inforsud-technologies.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Donne moi un byte slice de 99999 elements
	bs := make([]byte, 99999)
	// Body a une fonction read
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
```

L'interface Writer

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	io.Copy(os.Stdout, resp.Body)
}



```