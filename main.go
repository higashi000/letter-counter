package main

import (
	"github.com/higashi000/letter-counter/router"
)

func main() {
	lc := router.NewRouter()

	lc.E.Logger.Fatal(lc.E.Start(":5000"))
}
