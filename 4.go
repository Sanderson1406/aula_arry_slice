package main

import "fmt"

func main() {
	var num = [6]int{5, 8, 11, 18, 20, 25}
	soma := 0
	for i := 0; i < len(num); i++ {
		soma += num[i]
	}
	resul := soma / 6
	fmt.Print(resul)
}
