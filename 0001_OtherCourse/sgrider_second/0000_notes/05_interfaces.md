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