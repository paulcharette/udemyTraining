package main

import "fmt"

func main () {

	var large, small int;

	fmt.Print("Please enter a large number and a small number: ")
	fmt.Scan(&large, &small)
	fmt.Println("Remainder is ",large%small)
}
