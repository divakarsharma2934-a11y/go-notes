package main

import (
	"fmt"
	"time"
)

type coust struct {
	name  string
	phone string
}
type order struct {
	id     string
	price  float64
	status string
	date   time.Time
	coust
}

//
//func Neworder(id string, price float64, status string) *order {
//
//	myorder := order{
//		id:     id,
//		price:  price,
//		status: status,
//		//date:   time.Now(),
//	}
//	return &myorder
//}
//
//func (o *order) changed(status string) {
//	o.status = status
//}
//func (o order) getamount() float64 {
//	return o.price
//}

func main() {

	//newcoust := coust{
	//	name:  "divakar",
	//	phone: "13800138000",
	//}
	myoreder := order{
		id:     "1",
		price:  1.0,
		status: "pending",
		coust: coust{
			name:  "divakar",
			phone: "13800138000",
		},
	}
	myoreder.coust.name = "adity"
	fmt.Println(myoreder.coust)

	//
	//order1 := Neworder("2", 30, "confirmed")
	//fmt.Println(order1.getamount())

	//fmt.Println(myorder)
	//myorder.changed("confirmed")
	//fmt.Println(myorder)
	//fmt.Println(myorder.getamount())

	//lang := struct {
	//	language string
	//	isgood   bool
	//}{"Go", true}
	//fmt.Println(lang)

}
