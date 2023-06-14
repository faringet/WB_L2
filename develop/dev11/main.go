package main

import (
	"CitysTempRest/controllers"
	"CitysTempRest/initializers"
	"CitysTempRest/logging"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	logger := logging.GetLogger()
	logger.Info("Start app")
}

// @title           Temperature City Service
// @version         1.0
// @description     A weather service API in Go using Gin framework.

// @contact.name   Davydov Mikhail
// @contact.url    https://github.com/faringet/City_Temp_Rest_Api
// @contact.email  mik.davydov14@gmail.com

// @host      localhost:3000
// @BasePath  /subs
func main() {
	r := gin.Default()

	//TODO календарные api
	r.POST("/subs", controllers.SubsCreate)
	r.PUT("/subs/:id", controllers.SubsUpdate)
	r.GET("/subs", controllers.SubsIndex)
	r.GET("/subs/:id", controllers.SubsShow)
	r.DELETE("/subs/:id", controllers.SubsDelete)

	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
