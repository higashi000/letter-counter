package router

import (
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type LetterCounter struct {
	E *echo.Echo
}

type Msg struct {
	Text string `json:"text"`
}

func NewRouter() *LetterCounter {
	lc := &LetterCounter{
		E: echo.New(),
	}

	return lc
}

func CountLetter(text string) (int, int) {
	var availableSpace int
	var noSpace int

	noSpaceText := strings.Replace(text, " ", "", -1)

	noSpaceText = strings.Replace(noSpaceText, "　", "", -1)

	availableSpace = utf8.RuneCountInString(text)
	noSpace = utf8.RuneCountInString(noSpaceText)

	return availableSpace, noSpace
}

func (lc *LetterCounter) RecvText(c echo.Context) error {
	var msg Msg

	err := c.Bind(&msg)
	if err != nil {
		return errors.Wrap(err, "failed bind msg")
	}

	return nil
}
