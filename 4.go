package main

import "fmt"

func main() {
	var num = []float64{2.4, 13.14, 4.8, 1.8, 2.0, 2.5}
	var soma float64
	for i := 0; i < len(num); i++ {
		soma += num[i]
	}
	resul := soma / 6
	fmt.Print("A média é: ", resul)
}
