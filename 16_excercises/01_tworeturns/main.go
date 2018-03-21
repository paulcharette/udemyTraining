package main

import "fmt"

func halfAndEven (n int) (int,bool) {
	return n/2, n%2==0
}
func main() {
	fmt.Println(halfAndEven(1))
	fmt.Println(halfAndEven(76))
	fmt.Println(halfAndEven(13))

}
