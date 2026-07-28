import json

from aiokafka import AIOKafkaConsumer

from app.config.settings import KAFKA_BROKER, SENSOR_TOPIC
from app.models.sensor_event import SensorEvent
from app.services.analytics_service import process_sensor_event


async def consume_sensor_events():

    consumer = AIOKafkaConsumer(
        SENSOR_TOPIC,
        bootstrap_servers=KAFKA_BROKER,
        group_id="analytics-group",
        auto_offset_reset="earliest",
    )

    await consumer.start()

    print("✅ Kafka Consumer Connected")

    try:

        async for message in consumer:

            try:

                data = json.loads(message.value.decode())

                event = SensorEvent.model_validate(data)

                await process_sensor_event(event)

            except Exception as e:

                print("❌ Invalid Event:", e)

    finally:

        await consumer.stop()
        