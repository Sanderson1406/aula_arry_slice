package main

import "fmt"

func main() {
	nomes := []string{"Sanderson", "Maysa", "Eduarda", "Eric", "Luisa"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}
}
