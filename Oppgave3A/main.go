package main

import (
	"fmt"
)

func A(c1 chan int,c2 chan int) {

	fmt.Println("Enter digit")
	var x int
	fmt.Scanf("%d",&x)

	fmt.Println("Enter digit")
	var y int
	fmt.Scanf("%d",&y)

	c1 <- x  						// passing value to channel
	c1 <- y
	<- c2  							// receiving sum x+y from channel

}


func B (c1 chan int, c2 chan int){
	x,y := <-c1, <-c1 				//receiving digit x and digit y from channel
	c2 <- x+y         				//calculating x+y and passing the result to channel
}


func main() {
	c1,c2 := make(chan int), make(chan int)

	go A(c1,c2)
	go B(c1,c2)

	fmt.Println("The result of x + y is", <-c2)

}



