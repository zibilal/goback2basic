package router

import "github.com/gin-gonic/gin"

// "/ping" --> return pong

func SetupEngine(ping ...func(*gin.Context)) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping[0])
	r.GET("/ping/:name", ping[0])
	r.GET("/ping/:name/*payment", ping[1])
	// /ping/:name/lainnya/:apa
	return r
}
