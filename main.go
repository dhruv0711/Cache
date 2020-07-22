package main

import (
	_ "net/http/pprof"

	"github.com/dhruv0711/Cache/api"
	"github.com/dhruv0711/Cache/cache"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	cache.InitializeCache()

	r := gin.Default()
	pprof.Register(r)
	r.GET("/cache/fetch", api.Fetch)
	r.POST("/cache/insert", api.Insert)
	r.Run(":9091")
}
