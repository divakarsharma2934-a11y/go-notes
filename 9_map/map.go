package main

import (
	"fmt"
	"maps"
)

func main() {
	//m := make(map[string]int)
	//m["go"] = 1
	//fmt.Println(m)
	//fmt.Println(m["go"])

	//m := map[string]int{"sss": 1, "aaa": 2}
	//fmt.Println(m["ss"])
	//m["ss"] = 2
	//fmt.Println(m["ss"])

	m1 := map[string]int{"Phone": 2, "pc": 3}
	m2 := map[string]int{"Phone": 2}
	//a, ok := m1["Phon"]
	//fmt.Println(a)
	//if ok {
	//	fmt.Println("ok")
	//} else {
	//	fmt.Println("not ok")
	//}
	fmt.Println(maps.Equal(m1, m2))
}
