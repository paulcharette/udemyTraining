package main

import "fmt"

func maxNum (numSlice ...int) int {
	var max int
	for _,n := range numSlice {
		if n > max {
			max = n
		}
	}

	return max
}

func main () {
	fmt.Println(maxNum(10,200,13,23))
}
