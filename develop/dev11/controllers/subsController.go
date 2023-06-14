package controllers

import (
	"CitysTempRest/initializers"
	"CitysTempRest/logging"
	"CitysTempRest/models"
	"CitysTempRest/weather"
	"github.com/gin-gonic/gin"
)

// PostSub godoc
// @Summary Add a new city to a subscription
// @Description Takes a city JSON and store in DB. Return city's weather temperature.
// Produce json
// @Param city body models.Sub  true  "Sub JSON"
// @Success 200 {object} models.Sub
// @Router / [post]
func SubsCreate(c *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("SubsCreate worked")
	// Get data off req body
	var body struct {
		City        string
		Temperature float64
	}

	c.Bind(&body)

	// Create a sub
	sub := models.Sub{City: body.City}

	result := initializers.DB.Create(&sub)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"sub": sub,
	})

	//Put temperature to users City
	var a = weather.GetTemperature(weather.GetCityID(weather.ReturnMap(), sub.City))
	initializers.DB.First(&sub, sub.ID)
	initializers.DB.Model(&sub).Updates(models.Sub{Temperature: a})

}

// GetSub godoc
// @Summary Get all citys from subscription
// @Description Takes a citys and temperature JSON from DB. Return all city's weather temperature.
// Produce json
// @Success 200 {array} models.Sub
// @Router / [get]
func SubsIndex(c *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("SubsIndex worked")
	// Get the subs
	var subs []models.Sub
	initializers.DB.Find(&subs)

	// Respond with them
	c.JSON(200, gin.H{
		"sub": subs,
	})
}

// GetSub godoc
// @Summary Get single city from subscription
// @Description Takes a city and temperature JSON from DB. Return city weather temperature.
// Produce json
// @Param id path string true "search city by id"
// @Success 200 {object} models.Sub
// @Router /{id} [get]
func SubsShow(c *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("SubsShow worked")
	// Get id off url
	id := c.Param("id")

	// Get the subs
	var sub models.Sub
	initializers.DB.First(&sub, id)

	// Respond with them
	c.JSON(200, gin.H{
		"sub": sub,
	})
}

// PutSub godoc
// @Summary Update single city
// @Description Update single city.
// Produce json
// @Param id path string true "search city by id"
// @Param city body models.Sub  true  "Sub JSON"
// @Success 200 {object} models.Sub
// @Router /{id} [put]
func SubsUpdate(c *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("SubsUpdate worked")
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		City        string
		Temperature float64
	}

	c.Bind(&body)

	// Find the sub were updating
	var sub models.Sub
	initializers.DB.First(&sub, id)

	// Update it
	var a = weather.GetTemperature(weather.GetCityID(weather.ReturnMap(), sub.City))
	initializers.DB.Model(&sub).Updates(models.Sub{City: body.City, Temperature: a})

	// Respond with it
	c.JSON(200, gin.H{
		"sub": sub,
	})
}

// DeleteSub godoc
// @Summary Delete single city from subscription
// @Description Delete a city and temperature from DB.
// Produce json
// @Param id path string true "search city by id"
// @Success 200 {object} models.Sub
// @Router /{id} [delete]
func SubsDelete(c *gin.Context) {
	logger := logging.GetLogger()
	logger.Info("SubsDelete worked")
	// Get the id off the url
	id := c.Param("id")

	// Delete the subs
	initializers.DB.Delete(&models.Sub{}, id)

	// Respond
	c.Status(200)

}
