package handlers

import (
	"net/http"

	"device-service/internal/models"
	"device-service/internal/services"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"service": "Device Service",
		"status":  "Running",
	})
}

func HealthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

// POST /devices
func CreateDeviceHandler(c *gin.Context) {

	var device models.Device

	if err := c.ShouldBindJSON(&device); err != nil {

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Validation Failed",
		"error": err.Error(),
	})

	return
}

	createdDevice, err := services.CreateDevice(device)

	if err != nil {

	if err.Error() == "device already exists" {

		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	return
}

	c.JSON(http.StatusCreated, createdDevice)
}

// GET /devices/:id
func GetDeviceHandler(c *gin.Context) {

	deviceID := c.Param("id")

	device, err := services.GetDevice(deviceID)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Device not found",
		})
		return
	}

	c.JSON(http.StatusOK, device)
}

// GET /devices
func GetAllDevicesHandler(c *gin.Context) {

	devices, err := services.GetAllDevices()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// PUT /devices/:id
func UpdateDeviceHandler(c *gin.Context) {

	deviceID := c.Param("id")

	var device models.Device

	if err := c.ShouldBindJSON(&device); err != nil {

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Validation Failed",
		"error": err.Error(),
	})

	return
}

	device.DeviceID = deviceID

	updatedDevice, err := services.UpdateDevice(device)

	if err != nil {

	if err.Error() == "device not found" {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	return
}

	c.JSON(http.StatusOK, updatedDevice)
}

// DELETE /devices/:id
func DeleteDeviceHandler(c *gin.Context) {
    
	deviceID := c.Param("id")
    
	err := services.DeleteDevice(deviceID)

	if err != nil {

	if err.Error() == "device not found" {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	return
}

	c.JSON(http.StatusOK, gin.H{
		"message": "Device deleted successfully",
	})
}