package admin

import "gin"

func Index(c *gin.Context) {
	c.String(200,"Hello World !")
}