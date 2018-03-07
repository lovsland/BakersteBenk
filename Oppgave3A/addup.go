package main

import (
	"fmt"
	"time"
)

func main() {
	c:= make(chan int)
	go inputNumbers(c)
	time.Sleep(5*1e9)
	go addUp(c)
	time.Sleep(5*1e9)

	}

	func inputNumbers(c chan int){

		var x int
		var y int

		fmt.Println("Enter number:")
		fmt.Scan(&x)
		fmt.Println("Enter number:")
		fmt.Scan(&y)

		c<-x     				// passing value x through channel
		c<-y     				// passing value y through channel

		sum:= <- c  // receiving sum x + y from channel

		fmt.Println("The result of x + y is:", sum)
	}

	func addUp (c chan int){
		x, y := <-c, <-c  		// receiving value x and value y from channel
		sum := (x+y)

		c<- sum          		 // passing sum x+y through channel
	}








