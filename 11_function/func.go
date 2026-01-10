package main

import "fmt"

//func add(a, b int) int {
//	return a + b
//
//}

//func get() (string, string) {
//	return "go", "js"
//
//}

func to(fn func(a int) int) int {
	return fn(1)

}

func main() {
	//len1, len2 := get()
	//fmt.Println(len1, len2)
	fn := func(a int) int {
		return 2
	}
	aa := to(fn)
	fmt.Println(aa)

}
