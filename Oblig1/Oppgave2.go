package main
// Les https://en.wikipedia.org/wiki/Bubble_sort

import(
	"fmt"
	"math/rand"
	"time"

)
func main() {
	slice:=generateSlice(20)
	fmt.Println("\n---Unsorted---\n\n", slice)
	bubble_sort_modified(slice)
	fmt.Println("\n---Sorted---\n\n", slice, "\n")
	}
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int{

	slice:=make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
	}


func bubble_sort_modified(items []int){

	var (
		n = len(items)
		sorted = false
	)

for !sorted{
	swapped :=false
	for i:=0; i<n-1;i++ {
		if items[i] > items[i+1] {
			items[i+1], items[i] = items[i], items[i+1]
			swapped = true
		}
	}
	if !swapped {
		sorted = true}
		}
		n=n -1
	}





// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QuickSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}

