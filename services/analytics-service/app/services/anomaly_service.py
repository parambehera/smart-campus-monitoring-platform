from app.models.event_types import *


def detect_anomaly(sensor):

    anomalies = []

    if sensor.sensorType == "Temperature" and sensor.value > 45:
        anomalies.append(HIGH_TEMPERATURE)

    if sensor.sensorType == "Humidity" and sensor.value > 90:
        anomalies.append(HIGH_HUMIDITY)

    if sensor.sensorType == "Smoke" and sensor.value > 20:
        anomalies.append(SMOKE_DETECTED)

    if sensor.battery < 20:
        anomalies.append(LOW_BATTERY)

    return anomalies