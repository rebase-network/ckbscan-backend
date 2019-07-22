package main

import (
	"ckbscan-backend/conf"
	"ckbscan-backend/log"
	"ckbscan-backend/rest/block"
	"ckbscan-backend/rest/address"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

func main() {

	// TODO: used later, and need to add more functionality
	conf.Init()
	log.Init()
	//model.Init()

	startHttpSrv()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
}

func startHttpSrv() {
	g := gin.New()

	// Middlewares
	g.Use(gin.Recovery())

	g.GET("/api/getblock", rest.GetBlock)
	g.GET("/api/getaddress", address.GetAddress)

	g.Run()
}
