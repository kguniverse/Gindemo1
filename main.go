package main

import (
	"demo1/src/router"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func main() {
	r := gin.Default()
	router.CustomRoute(r)
	//model.Mysql()
	if err := r.Run(":8000"); err != nil {
		return
	}
}
