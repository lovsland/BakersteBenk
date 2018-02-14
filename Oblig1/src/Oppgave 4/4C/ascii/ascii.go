package ascii

import "fmt"


func ExtendedASCIIText() string {
	a1 := []byte("\x22\x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72\x22")
	for i := 0; i < len(a1); i++ {
		fmt.Printf("%c ", a1[i])
	}
	a2 := "\x22\x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72\x22"
	return a2
}
