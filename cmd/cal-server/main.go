package main

import (
	"fmt"
	"lightnet/pkg/calculator/handler"
	"lightnet/pkg/calculator/usecase"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//AUTH_KEY between cal-server and proxy-server need to be the same. This's because to prevent the traffic to go directly to cal-server.
	authKey := os.Getenv("AUTH_KEY")

	fmt.Println(authKey)

	e := echo.New()
	e.Use(middleware.Logger())

	calService := usecase.NewCalculatorService()
	handler.NewCalHandler(e, calService, authKey)

	e.Logger.Fatal(e.Start(":8888"))
}
