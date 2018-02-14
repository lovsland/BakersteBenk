package ascii

import (
	"testing"
	"fmt"
)

func TestExtendedASCIIText(t *testing.T) {
	ascii := ExtendedASCIIText()
	if isExtendedASCII(ascii) == false {
		t.Fail()
	}
}

func isExtendedASCII(a2 string) bool {
	for _, i := range a2 {
		if i < 0x80 || i > 0xff {
			return false
			fmt.Print("Inneholder andre tegn")
		}
	}
	fmt.Print("Inneholder bare extended ASCII")
	return true
}
