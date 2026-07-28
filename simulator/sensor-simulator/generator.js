import { devices } from "./devices.js";

function randomBattery() {
  return Math.floor(Math.random() * 41) + 60;
}

function randomTemperature() {
  return +(40 + Math.random() * 15).toFixed(2);
}

function randomHumidity() {
  return Math.floor(Math.random() * 41) + 40;
}

function randomSmoke() {
  return Math.floor(Math.random() * 16) + 5;
}

function randomMotion() {
  return Math.random() > 0.5;
}

export function generateReading() {
  const device =
    devices[Math.floor(Math.random() * devices.length)];

  let value;

  switch (device.sensorType) {
    case "Temperature":
      value = randomTemperature();
      break;

    case "Humidity":
      value = randomHumidity();
      break;

    case "Smoke":
      value = randomSmoke();
      break;

    case "Motion":
      value = randomMotion();
      break;

    default:
      value = 0;
  }

  return {
    deviceId: device.deviceId,
    sensorType: device.sensorType,
    building: device.building,
    value,
    unit: device.unit,
    battery: randomBattery(),
  };
}