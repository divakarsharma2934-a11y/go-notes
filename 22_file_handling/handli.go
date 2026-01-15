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

	//

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	file, err := dir.ReadDir(10)
	if err != nil {
		panic(err)
	}
	for _, file := range file {
		fmt.Println(file.Name(), file.IsDir())
	}
}
