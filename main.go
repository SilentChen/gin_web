package main

import (
	"fmt"
	"gin"
	"html/template"
	"io"
	"os"
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

	r.LoadHTMLGlob("views/**/*")

	r.Static("/static", "static")

	port := cfg.Key("app.port").String()

	r.Run(port)
}
