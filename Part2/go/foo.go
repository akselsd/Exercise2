// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	for {
		select {

		case update := <- add_number:
			i += update

		case signal := <- control:
			if signal == GetNumber{
				number <- i
			}

			if signal == Exit{
				return
			}
		}
	}
}

func incrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}

	finished <- true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}

	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	finished := make(chan bool, 2)
	number := make(chan int)
	add_number := make(chan int)
	control := make(chan int)

	go number_server(add_number, control, number)
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)

	<- finished
	<- finished

	control<-GetNumber
	Println("The magic number is:", <- number)
	control<-Exit
}
