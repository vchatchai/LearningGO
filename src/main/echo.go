package main

// import (
// 	"fizzbuzz"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// )

// // Handler
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

// func echoServer() {
// 	// Echo instance
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Routes
// 	e.GET("/:text", caesarHandler)

// 	e.GET("/fizzbuzz/:value", fizzbuzzHandler)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// // Handler
// func caesarHandler(c echo.Context) error {
// 	encrypted = "done"
// 	// encrypted := golang.Caesar(c.Param("text"), 4)
// 	// encrypted := golang.Caesar("Testted", 4)
// 	// println(c.Request.Param.GET)
// 	return c.String(http.StatusOK, encrypted) //map[string]string{"encrypted": encrypted})
// }

// // Handler
// func fizzbuzzHandler(c echo.Context) error {
// 	value, _ := strconv.Atoi(c.Param("value"))

// 	result := fizzbuzz.FizzbuzzCheck(value)
// 	// encrypted := golang.Caesar("Testted", 4)
// 	// println(c.Request.Param.GET)
// 	return c.String(http.StatusOK, result)
// }

// // func caesharH((c echo.Context)){

// // }
// // c caesharH((c echo.Context)){

// // }
