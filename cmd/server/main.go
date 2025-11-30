package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrbrist/poebin/internal/handler"
	"github.com/mrbrist/poebin/internal/middleware"
	"github.com/mrbrist/poebin/internal/r2"
	"github.com/mrbrist/poebin/web/layouts"
)

func main() {
	// Setup r2 integration
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Gin
	r := gin.Default()
	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &handler.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	r.Use(middleware.ErrorHandler())

	r.GET("/", func(c *gin.Context) {
		renderer := handler.New(c.Request.Context(), http.StatusOK, layouts.Home())
		c.Render(http.StatusOK, renderer)
	})

	r.GET("/:id", func(c *gin.Context) {
		renderer := handler.New(c.Request.Context(), http.StatusOK, layouts.Build())
		// id := c.Param("id")

		// build, err := r2.GetBuild(id)
		// if err != nil {
		// 	c.Error(err)
		//  return
		// }

		c.Render(http.StatusOK, renderer)
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
