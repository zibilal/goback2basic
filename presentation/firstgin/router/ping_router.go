package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Ping(c *gin.Context) {
	name := c.Param("name")
	if name != "" {
		c.String(200, fmt.Sprintf("pong %s", name))
		return
	}
	c.String(200, "pong")
}

func PingAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")

	str := fmt.Sprintf("%s the action %s", name, action)
	log.Println(str)

	c.String(200, fmt.Sprintf("pong %s", str))
}
