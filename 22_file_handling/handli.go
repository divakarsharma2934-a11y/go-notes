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

	//read
	//buffer := make([]byte, a)
	//d, err := file.Read(buffer)
	//if err != nil {
	//	panic(err)
	//}
	//for i := 0; i < len(buffer); i++ {
	//	println(d, string(buffer[i]))
	//}

	//

	//dir, err := os.Open(".")
	//if err != nil {
	//	panic(err)
	//}
	//defer dir.Close()
	//file, err := dir.ReadDir(10)
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range file {
	//	fmt.Println(file.Name(), file.IsDir())
	//}

	//write
	//f, err := os.Create("test.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()

	//f.WriteString("Hello World")
	//f.WriteString("\naaa a a")

	//a := []byte("hello world")
	//f.Write(a)

	//copy one file and add it content to other new file
	//file, err := os.Open("exa.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//file2, err := os.Create("exa2.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file2.Close()
	//reader := bufio.NewReader(file)
	//writer := bufio.NewWriter(file2)
	//
	//for {
	//	b, err := reader.ReadByte()
	//	if err != nil {
	//		if err.Error() != "EOF" {
	//			panic(err)
	//		}
	//		break
	//	}
	//	e := writer.WriteByte(b)
	//	if e != nil {
	//		panic(e)
	//	}
	//}
	//writer.Flush()
	//fmt.Println("return new to file")

	//delete a file
	err := os.Remove("exa2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("exa2.txt is removed")

}
