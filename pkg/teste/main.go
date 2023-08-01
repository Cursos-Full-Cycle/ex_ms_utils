package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4"}
	evento = evento[:2]
	fmt.Println(evento)
}
