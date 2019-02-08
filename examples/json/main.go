package main

import (
	"github.com/fatihkahveci/gin-inspector"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	debug := true

	if debug {
		r.Use(inspector.InspectorStats())
		r.GET("/_inspector", func(c *gin.Context) {
			c.JSON(200, inspector.GetPaginator())
		})
	}

	r.Run()
}
