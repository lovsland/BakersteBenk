package Oppgave_2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"log"
)

func main() {


	fileCount(getArg())

	r1 := runeCount(getArg())
	r2 := runeCount(getArg())
	var sorted []int = Bubble_sort_modified(r1)
	var slice []int = sorted[123:]
	getASCII(r2, slice)

}

func getArg() string {
	args := os.Args[1:]
	var st string

	for _, v := range args {
		st += v
	}
	return st
}

func fileCount(s string) {

	fmt.Println("Information about: test.txt ", s)
	fmt.Println("Number of lines in file: ", lineCount(s))

}

func lineCount(s string) int {
	fi, _ := os.Open("test.txt")
	fiScanner := bufio.NewScanner(fi)
	count := 0

	for fiScanner.Scan() {
		count++
	}
	return count
}

func runeCount(s string) []int {
	var countRune []int
	var count int
	data, err := ioutil.ReadFile("test.txt")

	for i := 0; i < 128; i++ {
		count = 0
		for _, v := range data {
			if err != nil {
				panic(err)
			}
			if int(v) == i {
				count++
			}
		}
		countRune = append(countRune, count)
	}
	return countRune
}

func checkChar() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)

	}
	count := strings.Count(string(data), "A")
	count2 := strings.Count(string(data), "1")
	fmt.Println(count, " ", count2)
}

func Bubble_sort_modified(list []int) []int {

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
	return list
}

func getASCII(data []int, sortedCount []int) {
	var hexa []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 128; j++ {
			if sortedCount[i] == data[j] {
				hexa = append(hexa, j)
			}
		}
	}
	fmt.Println("Most common runes: ")
	var place int = 1
	for i := 4; i >= 0; i-- {
		if hexa[i] == 32 {
			fmt.Println(place, ". Rune: space", ", ", "Counts: ", sortedCount[i])
			i -= 1
			place += 1
		}
		fmt.Println(place, ". Rune: ", string(hexa[i]), ", ", "Counts: ", sortedCount[i])
		place += 1
	}
}
