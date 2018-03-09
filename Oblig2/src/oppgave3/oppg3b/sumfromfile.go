package oppg3b

import (
	"strconv"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"time"
)


func ReadFromFile() int{
	file, err := ioutil.ReadFile(tekstfil)
	if err != nil {
		fmt.Printf("\n%s Kan ikke lese fil\n", tekstfil)
		time.Sleep(1*time.Second)
		os.Exit(1)
	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, "\n"))

	var ints [2]int

	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("Kunne ikke konvertere tallene til int")
			time.Sleep(1*time.Second)
			os.Exit(1)
		}
	}
	sum := ints[0] + ints[1]

	return sum
}

func WriteSum(sum int) {
	file, err := os.Create(tekstfil)
	if err != nil {
		fmt.Println("Feil ved oppretting av fil")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d", sum))
	if err != nil {
		fmt.Println("Kunne ikke skrive sum til fil")
		time.Sleep(1*time.Second)
		os.Exit(1)
	}
}
