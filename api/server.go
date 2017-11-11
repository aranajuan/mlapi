package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	startServer(true)
}

func startServer(debug bool) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	if debug {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	priceRoute(r)
	r.Run()
}
