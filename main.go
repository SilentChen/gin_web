package main

import (
	"gin"
	"web/libs"
	 _ "net/http/pprof"
)

func init() {

}

func main() {
	cfg := libs.LoadIniFile("conf/static.ini")

	if "dev" == cfg.Key("app.mode").String() {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	r := LoadRouters()

	r.LoadHTMLGlob("views/**/*")

	r.Static("/static", "static")

	port := cfg.Key("app.port").String()

	r.Run(port)
}
