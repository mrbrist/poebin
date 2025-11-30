package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrbrist/poebin/internal/r2"
)

func main() {
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		// id := c.Param("id")

		// build, err := r2.GetBuild(id)
		// if err != nil {
		// 	c.JSON(http.StatusNotFound, err)
		// }

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/:id/raw", func(c *gin.Context) {
		id := c.Param("id")

		build, err := r2.GetBuild(id)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}

		c.JSON(http.StatusOK, build.Data)
	})

	r.Run()
}
