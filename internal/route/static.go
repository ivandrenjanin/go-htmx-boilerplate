package route

import (
	"go-htmx/internal/locator"

	"github.com/gin-gonic/gin"
)

func StaticPublicHandlers(r *gin.Engine, locator locator.Locator) {
	staticPublic := r.Group("")

	staticPublic.GET("/", func(c *gin.Context) {
		c.HTML(200, "page/index.tmpl", gin.H{
			"title": "Hello, world!",
		})
	})

	staticPrivate := r.Group("/app")

	staticPrivate.GET("/", func(c *gin.Context) {
		c.HTML(200, "page/app.tmpl", gin.H{
			"title": "Hello, world!",
		})
	})
}
