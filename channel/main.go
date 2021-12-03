package main

import "fmt"

func main() {
	ch := make(chan int)
	go fib(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

func fib(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
	}

}
