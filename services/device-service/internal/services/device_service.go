package services

import (
	"errors"

	"device-service/internal/models"
	"device-service/internal/repository"
)

func CreateDevice(device models.Device) (models.Device, error) {

	// Check if device already exists
	_, err := repository.GetDevice(device.DeviceID)

	if err == nil {
		return models.Device{}, errors.New("device already exists")
	}

	err = repository.SaveDevice(device)

	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func GetDevice(deviceID string) (*models.Device, error) {

	return repository.GetDevice(deviceID)
}

func GetAllDevices() ([]models.Device, error) {

	return repository.GetAllDevices()
}

func UpdateDevice(device models.Device) (models.Device, error) {

	// Check whether device exists
	_, err := repository.GetDevice(device.DeviceID)

	if err != nil {
		return models.Device{}, errors.New("device not found")
	}

	err = repository.UpdateDevice(device)

	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func DeleteDevice(deviceID string) error {

	// Check whether device exists
	_, err := repository.GetDevice(deviceID)

	if err != nil {
		return errors.New("device not found")
	}

	return repository.DeleteDevice(deviceID)
}