package main

import (
	"fmt"
	"time"
)

func run(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 10; i++ {
		//go run(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)

}
