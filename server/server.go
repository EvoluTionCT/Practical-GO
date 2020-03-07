package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/photos", func(context echo.Context) error {
		pht := []byte(`{"title":"hello, it's me!!!"}`)
		return context.JSON(http.StatusOK,pht)
	})
	fmt.Println("Starting...")

	log.Fatal(e.Start(":1234"))
}
