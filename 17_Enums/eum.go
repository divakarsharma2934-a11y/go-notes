package main

import "fmt"

type orderstatus int

const (
	comfirmed orderstatus = iota
	received
	perpares
	delivered
)

func changestatus(status orderstatus) {
	fmt.Println("changing order status", status)
}

func main() {
	changestatus(comfirmed)

}
