package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// START OMIT

type Person struct {
	ID   string `uri:"id"   binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}

// END OMIT
