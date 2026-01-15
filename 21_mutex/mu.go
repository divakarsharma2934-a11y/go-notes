package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mux  sync.Mutex
}

func (p *post) incView(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mux.Unlock()
	}()
	p.mux.Lock()
	p.view += 1

}

func main() {
	var wg sync.WaitGroup
	mypost := post{view: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.incView(&wg)
	}
	wg.Wait()
	fmt.Println(mypost.view)

}
