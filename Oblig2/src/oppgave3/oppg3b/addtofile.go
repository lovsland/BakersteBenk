package oppg3b

import (
"fmt"
"os"
"io/ioutil"
"strconv"
"time"
)


var tekstfil = "3b.txt"

func Input() (int, int) {
	var tall1 int
	var tall2 int

	fmt.Println("Tall 1: ")
	_ , err := fmt.Scanf("%d", &tall1)
	if err != nil {
		fmt.Println("Må være en integer")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}

	fmt.Println("Tall 2: ")
	_, err = fmt.Scanf("%d", &tall2)
	if err != nil {
		_, err = fmt.Scanf("%d", &tall2)
		if err != nil {
			fmt.Println("Må være en integer")
			time.Sleep(1*time.Second)
			os.Exit(1)
		}
	}
	return tall1,tall2
}

func WriteNumbersToFile(tall1, tall2 int) {

	file, err := os.Create(tekstfil)
	if err != nil {
		fmt.Println("Feil ved oppretting av fil")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n%d" , tall1, tall2))
	if err != nil {
		fmt.Println("Kan ikke skrive tall til fil")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}
}

func PrintSum() {
	file, err := ioutil.ReadFile(tekstfil)
	if err != nil {
		fmt.Println("Kan ikke lese fra fil")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}

	fileString := string(file)
	sum, err := strconv.Atoi(fileString)
	if err != nil {
		fmt.Println("Kan ikke konvertere sum til int")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}

	fmt.Printf("Sum = %d\n", sum)
}
