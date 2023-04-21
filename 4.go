package main

import "fmt"

func main() {
	var num = [6]float64{5.8, 8.5, 11.8, 18.11, 20.18, 25.20}
	soma := 0
	for i := 0; i < len(num); i++ {
		soma += num[i]
	}
	resul := soma / 6
	fmt.Print(resul)
}
