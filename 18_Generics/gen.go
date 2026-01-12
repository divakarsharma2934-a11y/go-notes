package main

import "fmt"

func slice[T comparable](items []T) {
	for _, i := range items {
		fmt.Println(i)
	}
}

//type stack[T any] struct {
//	list []T
//}

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	//stts := []string{"a", "b", "c", "d", "e", "f"}
	boll := []bool{true, false}
	slice(boll)
	//mystack := stack[string]{
	//	list: []string{"gloang"},
	//}
	//fmt.Println(mystack)

}
