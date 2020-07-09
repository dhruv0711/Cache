package main

import (
	_ "net/http/pprof"

	"github.com/dhruv0711/Cache-Nokia/api"
	"github.com/dhruv0711/Cache-Nokia/cache"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	cache.InitializeCache()

	r := gin.Default()
	pprof.Register(r)
	r.GET("/fetch", api.Fetch)
	r.POST("/insert", api.Insert)
	r.Run(":9091")
}
