package api

import (
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)
	g.POST("/api/chart/add")
	g.POST("/api/chart/delete")
	g.POST("/api/chart/list")

	g.POST("/api/chart/build")
	g.POST("/api/chart/uninstall")
	g.POST("/api/chart/list")

	return g
}
