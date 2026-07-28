import json

from aiokafka import AIOKafkaProducer

from app.config.settings import (
    KAFKA_BROKER,
    ANALYTICS_TOPIC,
)

producer = None


async def start_producer():
    global producer

    producer = AIOKafkaProducer(
        bootstrap_servers=KAFKA_BROKER
    )

    await producer.start()

    print("✅ Analytics Producer Connected")
    print("Broker:", KAFKA_BROKER)
    print("Topic:", ANALYTICS_TOPIC)


async def stop_producer():
    global producer

    if producer:
        await producer.stop()


async def publish_event(event):
    global producer

    print("📤 About to publish:", event)

    try:
        await producer.send_and_wait(
            ANALYTICS_TOPIC,
            json.dumps(event).encode()
        )

        print("✅ Event Published Successfully")

    except Exception as e:
        print("❌ Publish Error:", e)



