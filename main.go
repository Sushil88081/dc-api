package main

import (
	"doctor-on-demand/config"

	"github.com/labstack/echo"
)

func main() {
	config.ConnectDB()
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
