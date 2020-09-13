package count

import (
	"strings"
	"unicode/utf8"

	"github.com/higashi000/letter-counter/router"
)

func CountLetter(text string) router.TextNum {
	var availableSpace int
	var noSpace int

	noSpaceText := strings.Replace(text, " ", "", -1)

	noSpaceText = strings.Replace(noSpaceText, "　", "", -1)

	availableSpace = utf8.RuneCountInString(text)
	noSpace = utf8.RuneCountInString(noSpaceText)

	return TextNum{availableSpace, noSpace}
}

func CountWord(text string) wordNum {

}
