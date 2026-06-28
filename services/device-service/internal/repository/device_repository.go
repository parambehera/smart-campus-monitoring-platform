package repository

import (
	"encoding/json"

	"device-service/internal/config"
	"device-service/internal/models"
)

const deviceSetKey = "devices"

// Save Device
func SaveDevice(device models.Device) error {

	deviceJSON, err := json.Marshal(device)

	if err != nil {
		return err
	}

	// Save Device
	err = config.RedisClient.Set(
		config.Ctx,
		device.DeviceID,
		deviceJSON,
		0,
	).Err()

	if err != nil {
		return err
	}

	// Save Device ID in Set
	return config.RedisClient.SAdd(
		config.Ctx,
		deviceSetKey,
		device.DeviceID,
	).Err()
}

// Get Single Device
func GetDevice(deviceID string) (*models.Device, error) {

	data, err := config.RedisClient.Get(
		config.Ctx,
		deviceID,
	).Result()

	if err != nil {
		return nil, err
	}

	var device models.Device

	err = json.Unmarshal([]byte(data), &device)

	if err != nil {
		return nil, err
	}

	return &device, nil
}

// Get All Devices
func GetAllDevices() ([]models.Device, error) {

	ids, err := config.RedisClient.SMembers(
		config.Ctx,
		deviceSetKey,
	).Result()

	if err != nil {
		return nil, err
	}

	var devices []models.Device

	for _, id := range ids {

		device, err := GetDevice(id)

		if err == nil {
			devices = append(devices, *device)
		}

	}

	return devices, nil
}

// Update Device
func UpdateDevice(device models.Device) error {

	deviceJSON, err := json.Marshal(device)

	if err != nil {
		return err
	}

	return config.RedisClient.Set(
		config.Ctx,
		device.DeviceID,
		deviceJSON,
		0,
	).Err()
}

// Delete Device
func DeleteDevice(deviceID string) error {

	err := config.RedisClient.Del(
		config.Ctx,
		deviceID,
	).Err()

	if err != nil {
		return err
	}

	return config.RedisClient.SRem(
		config.Ctx,
		deviceSetKey,
		deviceID,
	).Err()
}