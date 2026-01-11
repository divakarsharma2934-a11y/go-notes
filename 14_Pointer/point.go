package main

import "fmt"

func change(nums *int) {
	*nums = 5
	fmt.Println("changed no ", nums)

}

func main() {
	num := 1
	change(&num)
	fmt.Println("after change main", num)
}
