# Maps

## Qu'est ce que c'est ?

Il s'agit de la même chose qu'un dictionnaire en C#. Il s'agit d'un stockage de key/value pair dont les clés
sont d'un seul type et les values sont d'un seul type également.

```go
colors := map[string]string{
 	"red":   "#ff0000",
 	"green": "#4b34rd",
}
fmt.Println(colors)
// Second approach, empty map
var colors map[string]string

// last syntax, make tool
colors := make(map[string]string)

```

## Add and update item to map

```go
	colors := make(map[string]string)
	// var[key] = value
	colors["white"] = "#ffffff"
	fmt.Println(colors)

```

## Delete an item from the map

```go
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	fmt.Println(colors)
	delete(colors, "white")
	// Empty
	fmt.Println(colors)
```

## Iterate over a map

```go
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b34rd",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v has hex code %v\n", color, hex)
	}
}
```

## Quel est la différence entre les maps et les strucs

| Point       | Map                                                                                    | Struct                                                       |
| ----------- | -------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| Typage      | Toutes les clés doivent être du même type. Tous les valeurs doivent être du même type. | Les valeurs peuvent être de type différents                  |
| Utilisation | Utilisé pour représenter une collection de propriété liées                             | Utilisé pour représenter un objet avec différentes propriété |
| Type        | Type référence                                                                         | Type valeur                                                  |
| Indexation  | Les clés sont indexées, on peut itérer sur elles                                       | Les clés ne supportent pas l'indexation.                     |

L'avantage d'une map est également que vous n'avez pas besoin de connaitre toutes les clés à la compilation.
