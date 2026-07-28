import asyncio

from app.kafka.producer import start_producer
from app.consumers.sensor_consumer import consume_sensor_events


async def main():

    print("🚀 Analytics Service Started")

    await start_producer()

    await consume_sensor_events()


if __name__ == "__main__":
    asyncio.run(main())
    