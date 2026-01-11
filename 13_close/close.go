package main

import "fmt"

func counter() func() int {
	var count = 0
	return func() int {
		count += 1
		return count
	}

}

func main() {
	intt := counter()
	fmt.Println(intt())
	fmt.Println(intt())

}
