package controllers

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (server *Server) Home(c echo.Context) error {
	procc, err := strconv.Atoi(os.Getenv("MAX_PROC"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	runtime.GOMAXPROCS(procc)
	var messages = make(chan string)

	var print = func(who string) {
		var data = fmt.Sprintf("~Goroutine :  %s", who)
		messages <- data
	}

	go print("Hello, GoRest v01")

	var response = <-messages
	return c.JSON(http.StatusOK, response)
}
