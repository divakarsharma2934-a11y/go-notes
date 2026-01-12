package main

import "fmt"

type paymer interface {
	pay(amount float64)
}

type payment struct {
	gateway paymer
}

func (p payment) make(amonut float64) {
	p.gateway.pay(amonut)
}

type shipe struct {
}

func (o shipe) pay(a float64) {
	fmt.Println("shipe payment gatway", a)
}

type razor struct{}

func (r razor) pay(a float64) {
	fmt.Println("razor payment gatway", a)
}

type fake struct{}

func (f fake) pay(a float64) {
	fmt.Println("fack payment gatway", a)
}
func main() {
	//stripegat := shipe{}
	//razorgta := razor{}
	fakegta := fake{}
	newe := payment{
		gateway: fakegta,
	}
	newe.make(100)
}
