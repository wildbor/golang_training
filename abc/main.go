package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main(){
	e := echo.New()
	e.GET("/", HelloController)
	e.Start("8000")
}

func HelloController(c echo.Context) error{
	return c.String(http.StatusOK, "helloworld")
}
