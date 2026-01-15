package main

import (
	"fmt"
	"os"
)

func main() {
	//file, err := os.Open("exa.txt")
	//if err != nil {
	//	//error log
	//	panic(err)
	//}
	//fi, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(fi.Name())
	//a := fi.Size()
	//defer file.Close()
	//
	//buffer := make([]byte, a)
	//d, err := file.Read(buffer)
	//if err != nil {
	//	panic(err)
	//}
	//for i := 0; i < len(buffer); i++ {
	//	println(d, string(buffer[i]))
	//}

	data, err := os.ReadFile("exa.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
