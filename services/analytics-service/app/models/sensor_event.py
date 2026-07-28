from pydantic import BaseModel
from typing import Union

class SensorReading(BaseModel):

    deviceId: str

    sensorType: str

    building: str

    value: Union[float, bool]

    unit: str | None = None

    battery: int


class SensorEvent(BaseModel):

    eventType: str

    source: str

    timestamp: str

    payload: SensorReading