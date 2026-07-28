from dotenv import load_dotenv
import os

load_dotenv()

SERVICE_NAME = os.getenv("SERVICE_NAME")

KAFKA_BROKER = os.getenv("KAFKA_BROKER")

SENSOR_TOPIC = os.getenv("SENSOR_TOPIC")

ANALYTICS_TOPIC = os.getenv("ANALYTICS_TOPIC")