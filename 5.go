package main

import "fmt"

func main() {
	matriz := [3][4]int{{4, 5, 6, 7}, {7, 4, 1, 8}, {8, 5, 2, 3}}
	for linha := range matriz {
		for coluna := range matriz[linha] {
			matriz[linha][coluna] = linha + coluna
		}

	}
	fmt.Println(matriz)
}
