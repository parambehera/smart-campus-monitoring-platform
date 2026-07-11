import dotenv from "dotenv";

dotenv.config();

export const env = {
  PORT: process.env.PORT,
  SERVICE_NAME: process.env.SERVICE_NAME,
  KAFKA_BROKER: process.env.KAFKA_BROKER,
  KAFKA_TOPIC: process.env.KAFKA_TOPIC,
};