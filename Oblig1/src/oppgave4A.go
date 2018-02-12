package main

import (
	"fmt"
)

func main(){
	IterateOverExtendedASCIIStringLiteral()
}

func IterateOverExtendedASCIIStringLiteral() {
	// Kode for Oppgave 4a
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X %c %b \n", i, i, i)
	}
}