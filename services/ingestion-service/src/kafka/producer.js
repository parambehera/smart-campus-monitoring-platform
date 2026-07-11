import { Kafka } from "kafkajs";
import { env } from "../config/env.js";

const kafka = new Kafka({
  clientId: "ingestion-service",
  brokers: [env.KAFKA_BROKER],
});

export const producer = kafka.producer();

export async function connectProducer() {
  await producer.connect();
  console.log("✅ Kafka Producer Connected");
}

export async function publishEvent(event) {
  await producer.send({
    topic: env.KAFKA_TOPIC,
    messages: [
      {
        value: JSON.stringify(event),
      },
    ],
  });
}