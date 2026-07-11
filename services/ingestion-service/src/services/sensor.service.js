import { publishEvent } from "../kafka/producer.js";
import { env } from "../config/env.js";
import {EventTypes} from "./../kafka/eventTypes.js";
export async function ingestSensorData(sensorData) {
  const event = {
    eventType: EventTypes.SENSOR_READING_RECEIVED,
    source: env.SERVICE_NAME,
    timestamp: new Date(),
    payload: sensorData,
  };

  await publishEvent(event);

  return sensorData;
}