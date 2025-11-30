package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	templates "github.com/mrbrist/poebin/internal/handler"
	"github.com/mrbrist/poebin/internal/middleware"
	"github.com/mrbrist/poebin/internal/r2"
)

func main() {
	// Setup r2 integration
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Load templates
	err = templates.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	// Gin
	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := templates.Home(c.Writer, nil); err != nil {
			log.Println(err)
		}
	})

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		build, err := r2.GetBuild(id)
		if err != nil {
			c.Error(err)
			return
		}

		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := templates.Build(c.Writer, build); err != nil {
			log.Println(err)
		}
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
