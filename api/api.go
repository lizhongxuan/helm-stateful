package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"net/http"
)

func StartApiServer(port string) {
	// Create the Gin engine.
	g := gin.New()

	// Middlewares.
	g.Use()

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	pprof.Register(g)

	Load(
		g,
		gin.Recovery(),
	)

	log.Infof("Start to listen, port : %s", port)
	log.Info(http.ListenAndServe(port, g).Error())
}
