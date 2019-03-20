package main

import (
	"context"
	"fmt"
	"gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		"echo"	:	fmt.Sprintf,
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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown  Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
