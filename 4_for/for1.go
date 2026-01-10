package main

import "fmt"

//for is the only construct for looping

func main() {
	//while loop method
	//a := 0
	//for a <= 3 {
	//	fmt.Println(a)
	//	a = a + 1
	//}

	//infinte
	//for {
	//	fmt.Println(a)
	//}

	//basic for loop
	//for a := 0; a <= 3; a++ {
	//	fmt.Println(a)
	//
	//}

	//use of range
	for a := range 10 {
		fmt.Println(a)
	}
}
