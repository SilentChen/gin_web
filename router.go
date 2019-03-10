package main

import (
	"gin"
	"net/http/pprof"
	"web/controllers/admin"
	"web/controllers/api"
)

func ginPprof(router *gin.Engine) {

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

	// some handler init
	apiTest := new(api.Test)
	adminIndex := new(admin.Index)

	// link the route pattern to the handler

	r.GET("/api", apiTest.Index)
	r.GET("/admin", adminIndex.Index)
	ginPprof(r)
	/*articles := new(app.Articles)

	r.GET("/articles", articles.Index)
	r.GET("/article/create", articles.Create)
	r.GET("/article/edit/:id", articles.Edit)
	r.GET("/article/del/:id", articles.Del)
	r.POST("/article/store", articles.Store)*/

	return r
}