package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

// channel for SIGINT signal
var sigintCh = make(chan os.Signal,2)

func main() {

	signal.Notify(sigintCh, os.Interrupt, syscall.SIGINT)
	go stop()

	c:= make(chan int)
	go inputNumber(c)
	time.Sleep(5*1e9)
	go addUp(c)
	time.Sleep(5*1e9)

}


func inputNumber(c chan int){

	var x int
	var y int

	fmt.Println("Enter number:")
	_, err:= fmt.Scan(&x)

	if err !=nil {
		fmt.Printf("The number entered is invalid. ")
		os.Exit(1)

	}
	fmt.Println("Enter number:")
	_, err = fmt.Scan(&y)

	if err != nil {
		fmt.Printf("The number entered is invalid. ")
		os.Exit(1)
	}

	c<-x     	// passing value x through channel
	c<-y     	// passing value y through channel

	sum:= <- c  // receiving sum x + y from channel

	fmt.Println("The result of x + y is:", sum)
}


func addUp (c chan int){
	x, y := <-c, <-c  		// receiving value x and value y from channel
	sum := x+y

	c<- sum          		// passing sum x+y through channel
}


func stop() {
	<- sigintCh
	fmt.Println("The program was stopped before it finished")
	os.Exit(1)

}











