# Struct in go

## What is a struct

Collection de différentes propriété composant un type. Pour l'exemple des cartes, on aurait pu avoir une propriété "suit" et une "value".

### Créer une struct et une valeur

```go
// type NameOfTheStruct struct
type Person struct {
	firstname string
	lastname  string
}
func main() {
  // Create a new person. Un peu crade...
  alex := Person{ "Alex" , "Anderson" }
  // Il est mieux de préciser les fields
  alex := Person{firstname: "Alex", lastname: "Anderson"}
}
```

```go
import "fmt"

// Person struct
type Person struct {
	firstname string
	lastname  string
}

func main() {
  // Va initialiser les valeurs à leur zero value
	var alex Person
	fmt.Print(alex)
	fmt.Printf("%+v", alex)
}

```

### Update struct value

```go
import "fmt"

// Person struct
type Person struct {
	firstname string
	lastname  string
}

func main() {
	var alex Person
	alex.firstname = "Alex"
	alex.lastname = "Anderson"

	fmt.Print(alex)
	fmt.Printf("%+v", alex)
}
```

### Embedded structs

Il s'agit d'une sorte d'héritage. Cela évite la duplidation de code.

```go
package main

import "fmt"

// ContactInfo is
type ContactInfo struct {
	email   string
	zipCode int
}

// Person struct
type Person struct {
	firstname string
	lastname  string
  // Embedded struct. Possible de mettre simplement ContactInfo.
  //Cela déclarera une propriété ContactInfo du type ContactInfo
	contact ContactInfo
}

func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
	}}
	fmt.Printf("%+v", jim)
}
```

### Struct avec receiver

Il suffit de déclarer une fonction avec un receiver du type de la struct :)

```go
func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
	}}
	jim.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
```

!!!! Go passe les données par valeur et non par référence !!!!

```go
func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
	}}
	jim.updateName(("Jimmy"))
	jim.print()
}

// Ne mettra pas à jour "jim" en dehors de cette fonction. Pointer
func (p Person) updateName(newFirstName string) {
  //
	p.firstname = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
```

### Les pointers

Il est important de comprendre comment fonctionne la mémoire vive sur une machine.
Une représentation simple se ferait comme ceci. Des adresses dans la mémoire contiennent une valeur. Celle d'une variable.

| Address | Value                                     | var in code |
| ------- | ----------------------------------------- | ----------- |
| 0000    | Person{Jim Party {email@email.com 12000}} | ==> jim     |
| 0001    |                                           |             |
| 0002    |                                           |             |
| 0003    |                                           |             |
| 0004    |                                           |             |

Que ce passe t'il lorsque l'on appelle la méthode updateName ci-dessus

`jim.updateName(("Jimmy"))`

`func (p Person) updateName(newFirstName string)`

| Address | Value                                     | var in code |
| ------- | ----------------------------------------- | ----------- |
| 0000    | Person{Jim Party {email@email.com 12000}} | ==> jim     |
| 0001    |                                           |             |
| 0002    | Person{Jim Party {email@email.com 12000}} | ==> p       |
| 0003    |                                           |             |
| 0004    |                                           |             |

Go fait une copie de la valeur plutôt que d'utiliser la variable. On appelle cela le "Pass by value".

#### Comment contourner le problème ?

```go
func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
  }}
  // On récupère le pointer de jim en ajoutant &
	jimPointer := &jim
	jimPointer.updateName(("Jimmy"))
	jim.print()
}

// On indique que l'on va recevoir un pointer en ajoutant * au type du receiver. "Un pointer qui pointe vers une valeur de type Person
func (pointerToPerson *Person) updateName(newFirstName string) {
  // On va cherche la valeur de la variable du pointer en entourant
  // de parenthèse et en ajoutant * devant.
	(*pointerToPerson).firstname = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
```

En résumé:

**&** : Donne moi le pointer de cette variable

**\*** : Donne moi la valeur à cette adresse. Donne moi un accès direct à cette variable

| Address | Value                                     |
| ------- | ----------------------------------------- |
| 0000    | Person{Jim Party {email@email.com 12000}} |

Turn **Address** into a **Value** \*Value

Turn **Value** into a **Address** &Address

### Pointer shortcut

Il est possible de raccourcir l'implémentation ci-dessus.

```go
func main() {
	jim := Person{firstname: "Jim", lastname: "Party", contact: ContactInfo{
		email:   "email@email.com",
		zipCode: 12000,
	}}

	// Go est assez malin pour prendre le pointer si le receiver le demande
	// Il prendra le pointer automatiquement
	jim.updateName(("Jimmy"))
	jim.print()
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

```

### Pointer Gotchas

```go

func main(){
  mySlice := []string{"Hi", "There", "How"}
  updateSlice(mySlice)
  // mySlice sera mis à jour avec la valeur 0 à Bye...
  // Sans utiliser de pointer...
  fmt.Println(mySlice)
}

func updateSlice(s []string) {
  s[0] = "Bye"
}

```

### Reference vs Value Type

Dans le cas précédent, il faut comprendre comment go gère les slices.

Quant on créé un slice, go créé deux structures de données.

- un slice: il contient un pointer vers un tableau
- une capacité
- une longueur

En mémoire, cela sera stocké (à peu près) comme ça.

| Address | Value                        |
| ------- | ---------------------------- |
| 0000    | length, cap, pointer to head |
| 0001    | Actual array                 |

Lorsque l'on passe notre slice dans une fonction, on passe le slice qui pointe vers le tableau. On va donc modifier le slice en direct...

| Address | Value                        |
| ------- | ---------------------------- | ------------------- |
| 0000    | length, cap, pointer to head | slice dans main     |
| 0001    | Actual array                 | array des slices    |
| 0002    | length, cap, pointer to head | slice dans fonction |

Les types en go sont soit de type référence soit de type valeur

| Valeur (utiliser les pointer pour changer ces types dans une fonction ) | Reference (pas besoin de se soucier des pointer) |
| ----------------------------------------------------------------------- | ------------------------------------------------ |
| int                                                                     | slices                                           |
| float                                                                   | maps                                             |
| string                                                                  | channels                                         |
| bool                                                                    | pointers                                         |
| struct                                                                  | functions                                        |
