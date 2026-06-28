package models

type Device struct {
	DeviceID string `json:"deviceId" binding:"required"`
	Name string `json:"name" binding:"required"`
	SensorType string `json:"sensorType" binding:"required"`
	Building string `json:"building" binding:"required"`
	Floor int `json:"floor" binding:"gte=0"`
	Status string `json:"status" binding:"required,oneof=Active Inactive Maintenance"`
}