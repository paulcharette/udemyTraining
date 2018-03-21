package main

import "fmt"

func foo (...int) {
	fmt.Println("foo called")
}
func main () {
	fmt.Println("hey")
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
}
