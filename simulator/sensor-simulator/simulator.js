import axios from "axios";
import dotenv from "dotenv";
import { generateReading } from "./generator.js";

dotenv.config();

const URL = process.env.INGESTION_URL;
const INTERVAL = Number(process.env.INTERVAL);

console.log("🚀 Smart Campus Sensor Simulator Started");

setInterval(async () => {
  const reading = generateReading();
  try {
    await axios.post(URL, reading);
    
    console.log(
      `📤 ${reading.deviceId} (${reading.sensorType}) -> ${reading.value}${reading.unit}`
    );
  } catch (err) {
    console.log("❌", err.message);
  }
}, INTERVAL);