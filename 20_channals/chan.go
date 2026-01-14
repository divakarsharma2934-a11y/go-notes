package main

import "fmt"

// send
//func process(numchan chan int) {
//	for num := range numchan {
//		fmt.Println("processing no", num)
//		time.Sleep(time.Second)
//
//	}
//}

// receive
//func add(result chan int, a int, b int) {
//	numreslt := a + b
//	result <- numreslt
//}

//func task(done chan bool) {
//	defer func() { done <- true }()
//	fmt.Println("processing...")
//}

//func emailsend(emailchan chan string, done chan bool) {
//	defer func() { done <- true }()
//	for email := range emailchan {
//		fmt.Println("send email to", email)
//		time.Sleep(time.Second)
//	}
//}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case x := <-chan1:

			fmt.Println("val at chan1", x)
		case y := <-chan2:
			fmt.Println("val at chan2", y)
		}
	}
	//email := make(chan string, 100)
	//done := make(chan bool)
	//
	//go emailsend(email, done)
	//
	//for i := 0; i < 10; i++ {
	//	email <- fmt.Sprintf("%demail", i)
	//}
	//fmt.Println("done sending")
	//close(email)
	//<-done
	//email <- "1@email"
	//email <- "2@email"
	//email <- "3@email"
	//fmt.Println(<-email)
	//fmt.Println(<-email)
	//fmt.Println(<-email)

	//done := make(chan bool)
	//go task(done)
	//<-done

	//result := make(chan int)
	//
	//go add(result, 10, 20)
	//
	//myresult := <-result
	//fmt.Println(myresult)
	//numchan := make(chan int)
	//
	//go process(numchan)
	//
	//for {
	//	numchan <- rand.Intn(100)
	//}

}
