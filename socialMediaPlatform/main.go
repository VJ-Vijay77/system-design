package main

import (
	"socialMediaPlatform/common"
	"socialMediaPlatform/controllers"
	"socialMediaPlatform/database"
	"socialMediaPlatform/di"
	"socialMediaPlatform/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	common.LoadEnv()
	db := database.ConnectDB()
	common := common.Service{}
	controllers := &controllers.Service{
		Db:     db,
		Common: &common,
	}
	c := &di.Config{
		Config: controllers,
	}

	r := gin.Default()
	routes.InitRoutes(r, c)
	r.Run()
}
