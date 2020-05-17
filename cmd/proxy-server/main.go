package main

import (
	"net/url"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func noop(c echo.Context) (err error) {
	return err
}

func singleTargetBalancer(url *url.URL) middleware.ProxyBalancer {
	targets := []*middleware.ProxyTarget{
		{URL: url},
	}

	return middleware.NewRoundRobinBalancer(targets)
}

func main() {
	//AUTH_KEY between cal-server and proxy-server need to be the same. This's because to prevent the traffic to go directly to cal-server.
	authKey := os.Getenv("AUTH_KEY")

	origin, err := url.Parse("http://localhost:8888/")
	if err != nil {
		panic(err)
	}
	balancer := singleTargetBalancer(origin)

	addKeyMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Add("Authorization", "Bearer "+authKey)
			return next(c)
		}
	}

	e := echo.New()
	e.Use(addKeyMiddleware)
	e.Use(middleware.Logger())

	e.POST("/calculator.sum", noop, middleware.Proxy(balancer))
	e.POST("/calculator.sub", noop, middleware.Proxy(balancer))
	e.POST("/calculator.mul", noop, middleware.Proxy(balancer))
	e.POST("/calculator.div", noop, middleware.Proxy(balancer))

	e.Logger.Fatal(e.Start(":8000"))
}
