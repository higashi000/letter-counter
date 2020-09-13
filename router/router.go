package router

import (
	"fmt"
	"net/http"

	"github.com/higashi000/letter-counter/count"
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
	AvailableSpace int `json:"withspace"`
	NoSpace        int `json:"nospace"`
}

func NewRouter() *LetterCounter {
	lc := &LetterCounter{
		E: echo.New(),
	}

	lc.E.Use(middleware.Logger())
	lc.E.Use(middleware.Recover())

	lc.E.POST("/count", lc.RecvText)
	lc.E.GET("/", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "./page/index.html")

		return nil
	})
	lc.E.GET("/js", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "./page/index.js")

		return nil
	})

	return lc
}

func (lc *LetterCounter) RecvText(c echo.Context) error {
	var msg Msg

	err := c.Bind(&msg)
	if err != nil {
		return errors.Wrap(err, "failed bind msg")
	}

	fmt.Println(msg.Text)
	textnum := count.CountLetter(msg.Text)

	fmt.Println(textnum.AvailableSpace)

	c.JSON(http.StatusOK, textnum)

	return nil
}
