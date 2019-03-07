package main

import (
	"fmt"
	"gin"
	"io"
	"os"
	"web/controllers/api"
	"web/libs"
)

func LoadRouters() *gin.Engine {
	// log path setting
	webRootPath, err := os.Getwd()
	libs.CheckErr(err)
	logPath := webRootPath + "/logs/access.log"
	logFile, _ := os.Create(logPath)

	// request log out put, file and terminate stdout
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	r := gin.New()

	// setting the logger format and use gin logger module
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// use gin recovery module
	r.Use(gin.Recovery())

	// some handler init
	apiTest := new(api.Test)

	// link the route pattern to the handler

	r.GET("/api", apiTest.Index)

	/*articles := new(app.Articles)

	r.GET("/articles", articles.Index)
	r.GET("/article/create", articles.Create)
	r.GET("/article/edit/:id", articles.Edit)
	r.GET("/article/del/:id", articles.Del)
	r.POST("/article/store", articles.Store)*/

	return r
}