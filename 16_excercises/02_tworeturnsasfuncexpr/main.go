package main

import "fmt"


func main() {

	halfAndEven := func (n int) (int,bool) {
		return n/2, n%2==0
	}

	fmt.Println(halfAndEven(1))
	fmt.Println(halfAndEven(76))
	fmt.Println(halfAndEven(13))

}
