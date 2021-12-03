package main

import "fmt"

func main() {
	ch := make(chan int)
	qCh := make(chan struct{})
	go fib(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	qCh <- struct{}{}
	<-qCh
	fmt.Println("end")

}

func fib(ch chan int, qCh chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			qCh <- struct{}{}
			return
		}

	}

}
