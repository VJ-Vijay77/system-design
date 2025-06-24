package main

import (
	"socialmedia/common"
	"socialmedia/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	common.LoadEnv()
	service := common.GetServiceContext()
	r := gin.Default()
	routes.InitRoutes(r, service)
	r.Run()
}
