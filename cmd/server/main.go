package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrbrist/poebin/internal/r2"
	"github.com/mrbrist/poebin/middleware"
)

func main() {
	// Setup r2 integration
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Gin
	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	r.GET("/:id", func(c *gin.Context) {
		// id := c.Param("id")

		// build, err := r2.GetBuild(id)
		// if err != nil {
		// 	c.Error(err)
		//  return
		// }

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/:id/raw", func(c *gin.Context) {
		id := c.Param("id")

		build, err := r2.GetBuild(id)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, build)
	})

	r.Run()
}
