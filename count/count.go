package count

import (
	"strings"
	"unicode/utf8"
)

func CountLetter(text string) TextNum {
	var availableSpace int
	var noSpace int

	noSpaceText := strings.Replace(text, " ", "", -1)

	noSpaceText = strings.Replace(noSpaceText, "　", "", -1)

	availableSpace = utf8.RuneCountInString(text)
	noSpace = utf8.RuneCountInString(noSpaceText)

	return TextNum{availableSpace, noSpace}
}
