package ascii

import "fmt"

//Kode for oppgave 4a
func IterateOverASCIIStringLiteral(sl string) {
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%X %c %b\n", sl[i], sl[i], sl[i])
	}
}

// Kode for Oppgave 4b
func ExtendedASCIIText() {
	bytes := []byte("\x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%c ", bytes[i])
	}
}