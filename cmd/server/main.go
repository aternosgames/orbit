package main

import "github.com/labstack/echo"

func main() {
	echoServer := echo.New()

	echoServer.GET("/", func(context echo.Context) error {
		context.JSON(200, "Simple HTTP Server")

		return nil
	})
}
