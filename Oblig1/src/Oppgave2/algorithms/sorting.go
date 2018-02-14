package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {

	N := len(list)
	var i int

	for i = 0; i < N; i++ {
	sweep(list, i)

	}
	}
	func sweep(numbers []int, prevPasses int) {
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
