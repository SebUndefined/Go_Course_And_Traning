# Basics

## 5 questions importantes

**Comment on lance un programme go ?**

Il suffit de taper *go run* suivi du nom du fichier à lancer (go run  main.go). 
Il existe d'autres commande:
- **go run:** compile et exécute un ou deux fichiers
- **go build:** Compile le programme
- **go fmt:** formate le fichier
- **go install:** compile et installe un package
- **go get : dwl** raw src code de quelquq'un d'autre
- **go test: lance** les tests du code

**Que veut dire le package main ?**

Package == Project == Workspace
Il y a deux types de package
- Executable: Du code que l'on a tapé et destiné à être exécuté
- Réutilisable: Code utilisé comme 'helper', 'librairie'. Du code destiné à être intégré à un code exécutable. 

Le nom du package détermine si c'est un code exécutable ou réutilisable. main est utilisé pour construire un programme exécutable. Si on utilise un autre nom, le build ne va pas construire un exécutable mais un code utilisable comme dépendance (helper). 

**Qu'est ce que l'import fmt ?**

fmt est utilisé comme import. Un import est utilisé pour intégrer et utiliser du code d'un package externe. Il peut s'agir de lib standard ou de lib développée par d'autres ingénieurs. 

**Qu'est ce que 'func' ?**

func est là pour déclarer une fonction. 

**Comment le fichier main.go est organisé ?**
1. Déclaration du package
2. Import: ce dont on a besoin pour le fichier courant
3. body: logique qui fait quelque chose. 


```go
package main

import "fmt"

func main() {
	fmt.Println("Hi there")
}

```