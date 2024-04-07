package main

import "github.com/gin-gonic/gin"

// START OMIT

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/products", legacyGetProduct)
		v1.POST("/products", legayPostProduct)
		v1.PUT("/products", legacyPutProduct)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/products", newGetProduct)
		v2.POST("/products", newPostProduct)
		v2.PUT("/products", newPutProduct)
	}

	router.Run(":8080")
}

// END OMIT
