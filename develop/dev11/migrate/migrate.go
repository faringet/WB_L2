package main

import (
	"CitysTempRest/initializers"
	"CitysTempRest/logging"
	"CitysTempRest/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Sub{})
	logger := logging.GetLogger()
	logger.Info("Start migration")
}
