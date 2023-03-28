package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
	num = append(num[:3], num[4:]...)
	fmt.Println(num)
}
