from datetime import datetime

from app.config.settings import SERVICE_NAME
from app.kafka.producer import publish_event

from app.models.sensor_event import SensorEvent
from app.services.statistics_service import update_statistics
from app.services.anomaly_service import detect_anomaly


async def process_sensor_event(event: SensorEvent):

    # Update statistics for numeric sensors
    if not isinstance(event.payload.value, bool):

        stats = update_statistics(
            event.payload.deviceId,
            event.payload.value
        )

        print("\n========== LIVE ANALYTICS ==========")
        print(f"Device      : {stats['deviceId']}")
        print(f"Current     : {stats['current']}")
        print(f"Average     : {stats['average']}")
        print(f"Minimum     : {stats['minimum']}")
        print(f"Maximum     : {stats['maximum']}")
        print(f"Readings    : {stats['count']}")
        print("====================================")

    # Detect anomalies
    anomalies = detect_anomaly(event.payload)

    # Publish every anomaly to Kafka
    for anomaly in anomalies:

        analytics_event = {
            "eventType": anomaly,
            "source": SERVICE_NAME,
            "timestamp": datetime.now().isoformat(),
            "payload": event.payload.model_dump()
        }

        print("📤 Publishing:", analytics_event)

        await publish_event(analytics_event)

        print(f"✅ Published: {anomaly}")


