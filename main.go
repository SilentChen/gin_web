package main

import (
	"gin"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
	"web/libs"
)

func init() {
	// log path setting
	webRootPath, err := os.Getwd()
	libs.CheckErr(err)
	logPath := webRootPath + "/logs/access.log"
	logFile, _ := os.Create(logPath)

	// request log out put, file and terminate stdout
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func main() {
	cfg := libs.LoadIniFile("conf/static.ini")

	if "dev" == cfg.Key("app.mode").String() {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	r := LoadRouters()

	r.Delims("{{", "}}")

	r.SetFuncMap(template.FuncMap{

	})

	r.LoadHTMLGlob("views/**/*")

	r.Static("/static", "static")

	port := cfg.Key("http.port").String()

	s := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
