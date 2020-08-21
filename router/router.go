package router

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

type LetterCounter struct {
	E *echo.Echo
}

type Msg struct {
	Text string `json:"text"`
}

type TextNum struct {
	AvailableSpace int `json:"availablespace"`
	NoSpace        int `json:"nospace"`
}

func NewRouter() *LetterCounter {
	lc := &LetterCounter{
		E: echo.New(),
	}

	lc.E.Use(middleware.Logger())
	lc.E.Use(middleware.Recover())

	lc.E.GET("/count", lc.RecvText)

	return lc
}

func CountLetter(text string) TextNum {
	var availableSpace int
	var noSpace int

	noSpaceText := strings.Replace(text, " ", "", -1)

	noSpaceText = strings.Replace(noSpaceText, "　", "", -1)

	availableSpace = utf8.RuneCountInString(text)
	noSpace = utf8.RuneCountInString(noSpaceText)

	return TextNum{availableSpace, noSpace}
}

func (lc *LetterCounter) RecvText(c echo.Context) error {
	var msg Msg

	err := c.Bind(&msg)
	if err != nil {
		return errors.Wrap(err, "failed bind msg")
	}

	fmt.Println(msg.Text)
	textnum := CountLetter(msg.Text)

	fmt.Println(textnum.AvailableSpace)

	c.JSON(http.StatusOK, textnum)

	return nil
}
