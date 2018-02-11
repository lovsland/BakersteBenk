package main

import "fmt"

func main() {
	var list[]int = []int{5,4,2,3,1,0}
	fmt.Println("Our list of numbers is:", list)
	Bubble_sort_modified(list)
	fmt.Println("After sorting:", list)
}
// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {

	N:=len(list)
	var i int

	for i =0; i<N; i++ {
		sweep(list,i)

	}
}
func sweep(numbers[]int, prevPasses int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
