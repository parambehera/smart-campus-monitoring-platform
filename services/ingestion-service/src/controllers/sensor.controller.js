import { ingestSensorData } from "../services/sensor.service.js";

export async function receiveSensorData(req, res) {
  try {
    const sensorData = req.body;

    await ingestSensorData(sensorData);

    return res.status(201).json({
      message: "Sensor data received successfully",
      data: sensorData,
    });
  } catch (err) {
    return res.status(500).json({
      error: err.message,
    });
  }
}