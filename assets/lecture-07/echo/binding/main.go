package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// START OMIT

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		executeSomeBusinessLogic(u)
		return c.JSON(http.StatusOK, u)
	})
}

// END OMIT
