package main

import "fmt"

func main() {
	//nums := []int{12, 2, 3, 34, 5}
	//sum := 0
	//for i, n := range nums {
	//	fmt.Println(n, i)
	//	sum += n
	//}
	//fmt.Println(sum)

	m := map[string]string{"fname": "Divakar", "lname": "Sharma"}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
