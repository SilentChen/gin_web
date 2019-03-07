package api

import "gin"

type Test struct {
	Base
}

func (_ *Test)Index(c *gin.Context)  {
	c.String(200,"Hello World ! ")
}