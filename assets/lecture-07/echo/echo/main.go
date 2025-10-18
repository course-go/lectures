package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// START OMIT

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Get("/user/:id", getUser)
	e.Get("/show", show)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// END OMIT
