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
	db.GetRow("select * from game_roles limit 1", record)
	log.Println(record)

	// var records []map[string]string
	// db.GetAll("select * from game_roles where accname = 'test'", &records)
	// log.Println(records)

	c.JSON(200,gin.H{
		"code"		:	200,
		"message"	:	"success",
	})
}
