package main

import (
	"fmt"
	"gin"
	"net/http/pprof"
	_"net/http/pprof"
	"web/controllers/admin"
	"web/controllers/api"
)

func profRouter(router *gin.Engine) {

	router.GET("/debug/pprof", func(c *gin.Context) {
		pprof.Index(c.Writer, c.Request)
	})

	router.GET("/debug/pprof/heap", func(c *gin.Context) {
		pprof.Handler("heap").ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/block", func(c *gin.Context) {
		pprof.Handler("block").ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/goroutine", func(c *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/threadcreate", func(c *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/cmdline", func(c *gin.Context) {
		pprof.Cmdline(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/profile", func(c *gin.Context) {
		pprof.Profile(c.Writer, c.Request)
	})
	router.GNP("/debug/pprof/symbol", func(c *gin.Context) {
		pprof.Symbol(c.Writer, c.Request)
	})
	router.GET("debug/pprof/trace", func(c *gin.Context) {
		pprof.Trace(c.Writer, c.Request)
	})
	router.GET("/debug/pprof/mutex", func(c *gin.Context) {
		pprof.Handler("mutex").ServeHTTP(c.Writer, c.Request)
	})
}

func LoadRouters() *gin.Engine {

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
	adminIndex := new(admin.Index)

	// link the route pattern to the handler
	r.GET("/api", apiTest.Index)
	r.GET("/admin", adminIndex.Index)
	profRouter(r)
	/*articles := new(app.Articles)

	r.GET("/articles", articles.Index)
	r.GET("/article/create", articles.Create)
	r.GET("/article/edit/:id", articles.Edit)
	r.GET("/article/del/:id", articles.Del)
	r.POST("/article/store", articles.Store)*/

	return r
}