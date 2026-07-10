package services

import (
	"errors"
	"time"

	"device-service/internal/config"
	"device-service/internal/kafka"
	"device-service/internal/models"
	"device-service/internal/repository"
)

func CreateDevice(device models.Device) (models.Device, error) {

	// Check if device already exists
	_, err := repository.GetDevice(device.DeviceID)

	if err == nil {
		return models.Device{}, errors.New("device already exists")
	}

	// Save device in Redis
	err = repository.SaveDevice(device)

	if err != nil {
		return models.Device{}, err
	}

	// Create Kafka Event
	event := kafka.Event{
		EventType: kafka.DeviceCreated,
		Source:    config.GetEnv("SERVICE_NAME"),
		Timestamp: time.Now(),
		Payload:   device,
	}

	// Publish Event
	err = kafka.PublishEvent(event)

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

	// Update in Redis
	err = repository.UpdateDevice(device)

	if err != nil {
		return models.Device{}, err
	}

	// Create Kafka Event
	event := kafka.Event{
		EventType: kafka.DeviceUpdated,
		Source:    config.GetEnv("SERVICE_NAME"),
		Timestamp: time.Now(),
		Payload:   device,
	}

	// Publish Event
	err = kafka.PublishEvent(event)

	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func DeleteDevice(deviceID string) error {

	// Check whether device exists
	device, err := repository.GetDevice(deviceID)

	if err != nil {
		return errors.New("device not found")
	}

	// Delete from Redis
	err = repository.DeleteDevice(deviceID)

	if err != nil {
		return err
	}

	// Create Kafka Event
	event := kafka.Event{
		EventType: kafka.DeviceDeleted,
		Source:    config.GetEnv("SERVICE_NAME"),
		Timestamp: time.Now(),
		Payload:   device,
	}

	// Publish Event
	err = kafka.PublishEvent(event)

	if err != nil {
		return err
	}

	return nil
}