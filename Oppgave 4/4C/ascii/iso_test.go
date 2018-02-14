package ascii

import "testing"

func TestExtendedASCIIText(t *testing.T) {
	ascii := ExtendedASCIIText()
	if isExtendedASCII(ascii) == false {
		t.Fail()
	}
}

func isExtendedASCII(a2 string) bool {
	for _, i := range a2 {
		if i > 127 || i <= 255 {
			return false
		}
	}
	return true
}
