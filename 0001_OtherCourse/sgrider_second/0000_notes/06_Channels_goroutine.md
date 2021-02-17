# Channel and Goroutine

Pour illustrer les channels et go routines, on va faire un site checker. Cela consiste à checker si un site est en ligne ou pas. Il devra analyser plusieurs sites en même temps. Nous allons faire d'abord la méthode naïve puis on va utiliser des goroutines et channels.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Déclaration des liens à surveiller
	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amason.com",
	}
	// Loop sur la liste pour checker les sites les uns après les autres
	for _, link := range links {
		// On check le lien
		checkLink(link)
	}
}

func checkLink(link string) {
	// On ne se souci pas de resp ("_") et on prend simplement l'erreur
	_, err := http.Get(link)
	// S'il y a une erreur, on affiche que le site est probablement down
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}
	// Sinon, on affiche que le site est up
	fmt.Println(link, " is up")
}
```

l'output est la suivante:

```
$ go run main.go 
https://www.facebook.com  is up
https://www.google.com  is up
https://www.stackoverflow.com  is up
https://www.golang.org  is up
https://www.amason.com  might be down
```

On peut voir que:
- les sites sont analysés dans l'ordre de la liste
- qu'il y a un délais entre les sites
- si on veut rechecker un site, il faut attendre que la liste complète soit terminée pour la recommencer

Cela veut simplement dire que le programme check les sites les uns après les autres. La fonction n'avance pas tant que le site n'a pas été checké (réponse reçu). Il s'agit d'une fonction bloquante. Pour 5 sites, cela ne pose pas de problème. Cela peut le devenir si nous avons plusieurs centaines de sites. 

On pourrait avoir un check parallèle (plutôt qu'en série), en lançant les checks en même temps et attendre le résultat de ces derniers. C'est là que les goroutines et les channels entrent en jeux !!!

## Théorie des go routine

Première chose à savoir, notre programme est une go routine en elle même !!

L'idée est de résoudre le problème des fonctions bloquante (ici http.Get) en les lançant dans une autre go routine. Simplement en ajoutant go devant l'appel de la fonction... Simple is that !!

Les goroutines sont lancées par le go Scheduler. Il se charge de répartir les goroutines sur les coeurs du cpu. Si on a un cpu, on ne pourra lancer que une goroutine à la fois, si on en a deux, deux goroutines etc.

## Concurrence vs parralèle:

La concurrence n'est pas le parralélisme

### Concurrence

On peut avoir plusieurs threads qui exécute du code. Si un thread est bloquant, un autre est alors pris et exécuté. Il ne s'agit pas d'éxécuter des threads en même temps, mais de les gérer en même temps. On les programme et on change de threads en fonction des blocages de ces derniers.

### Parralélisme

Plusieurs threads exécutant au même moment (exactement) requiert plusieurs CPU. On peut faire plusieurs choses en même temps.

## Implémentation des goroutines

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Déclaration des liens à surveiller
	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amason.com",
	}
	// Loop sur la liste pour checker les sites les uns après les autres
	for _, link := range links {
		// On check le lien en la lançant dans une goroutine
		go checkLink(link)
	}
}

func checkLink(link string) {
	// On ne se souci pas de resp ("_") et on prend simplement l'erreur
	_, err := http.Get(link)
	// S'il y a une erreur, on affiche que le site est probablement down
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}
	// Sinon, on affiche que le site est up
	fmt.Println(link, " is up")
}
```

Le problème si l'on ne fait que cela est que la sortie du programme ne donnera rien dans la plupart des cas. En effet, le programme lance la main goroutine et les childs goroutines, la main goroutine se termine sans attendre le résultat des childs car la main goroutine n'a plus rien à faire... Il ne suffit pas simplement de mettre go devant l'appel de fonction, il faut utiliser des channels afin d'informer la main goroutine des child goroutine. Ces channels permettent la communication entre les goroutines. Ici, entre la main goroutine et les childs. 

Un channel est typé ou plutôt la valeur que transporte le channel est typée.



### Etape 1 déclarer le channel

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Déclaration des liens à surveiller
	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amason.com",
	}
	// Créer un channel de type string
	c := make(chan string)
	// Loop sur la liste pour checker les sites les uns après les autres
	for _, link := range links {
		// On check le lien et on passe le channel en paramètre
		go checkLink(link, c)
	}
}

func checkLink(link string, c chan string) {
	// On ne se souci pas de resp ("_") et on prend simplement l'erreur
	_, err := http.Get(link)
	// S'il y a une erreur, on affiche que le site est probablement down
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}
	// Sinon, on affiche que le site est up
	fmt.Println(link, " is up")
}
```

Le channel est donc disponible dans notre fonction. Comment envoyer des données à ce channel ? 

```go
// Envoyer la valeur 5 dans le channel
channel <- 5
// Attendre qu'une valeur soit dans le channel. Quand cela est le cas,
// assigner cette valeur à "myNumber"
myNumber <- channel
// Attendre qu'une valeur soit dans le channel. Quand cela est le cas,
// logger cette valeur immédiatement
fmt.Println(<- channel)
```

```go
package main

import (
	"fmt"
	"net/http"
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
	// Attendre une valeur et la logger immédiatement
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	// On ne se souci pas de resp ("_") et on prend simplement l'erreur
	_, err := http.Get(link)
	// S'il y a une erreur, on affiche que le site est probablement down
	if err != nil {
		fmt.Println(link, " might be down")
		// Envoyer un message au channel
		c <- "Might be down I think"
		return
	}
	// Sinon, on affiche que le site est up
	fmt.Println(link, " is up")
	// Envoyer un message au channel
	c <- "Yep it's up !"
}
```

Le problème avec ce programme est que l'on attend 