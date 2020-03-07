package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvoluTionCT/Practical-GO/typicode"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/photos", func(context echo.Context) error {
		tc := typicode.New("/photos")
		p := []typicode.Photo{}
		err := tc.Get(&p)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err)
		}
		return context.JSON(http.StatusOK, p)
	})
	fmt.Println("Starting...")

	println()
	log.Fatal(e.Start(":1234"))

	recover()
}

func Go(fn func()) {
	go func() {

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("go-routine panic: %v\n%s", err, buf)
			}
		}()

		fn()
	}()
}
