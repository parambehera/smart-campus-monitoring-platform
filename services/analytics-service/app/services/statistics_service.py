from datetime import datetime

statistics = {}


def update_statistics(device_id, value):

    if device_id not in statistics:

        statistics[device_id] = {
            "count": 0,
            "sum": 0,
            "min": value,
            "max": value,
            "current": value,
            "lastUpdated": datetime.now().isoformat()
        }

    stats = statistics[device_id]

    stats["count"] += 1
    stats["sum"] += value

    stats["min"] = min(stats["min"], value)
    stats["max"] = max(stats["max"], value)

    stats["current"] = value

    stats["lastUpdated"] = datetime.now().isoformat()

    average = stats["sum"] / stats["count"]

    return {
        "deviceId": device_id,
        "current": stats["current"],
        "average": round(average, 2),
        "minimum": stats["min"],
        "maximum": stats["max"],
        "count": stats["count"],
        "lastUpdated": stats["lastUpdated"]
    }