package api

import (
	"gin"
	"log"
	"web/models"
)

type Test struct {
	Base
}

func (_ *Test)Index(c *gin.Context)  {
	db := new(models.Mysql)
	record := make(map[string]string)
	db.GetAll(record)
	log.Println(record)
	c.String(200,"Hello World ! ")
}