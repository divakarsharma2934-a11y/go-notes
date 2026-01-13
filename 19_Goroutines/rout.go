package main

import (
	"fmt"
	"sync"
)

func run(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go run(i, &wg)
		//go func(i int) {
		//	fmt.Println(i)
		//}(i)
	}
	//time.Sleep(1 * time.Second)
	wg.Wait()

}
