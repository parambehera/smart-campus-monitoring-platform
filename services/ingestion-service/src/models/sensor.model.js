export class SensorReading {
  constructor(
    deviceId,
    sensorType,
    building,
    value,
    unit,
    battery
  ) {
    this.deviceId = deviceId;
    this.sensorType = sensorType;
    this.building = building;
    this.value = value;
    this.unit = unit;
    this.battery = battery;
  }
}