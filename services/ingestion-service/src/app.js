import express from "express";
import sensorRoutes from "./routes/sensor.routes.js";

const app = express();

app.use(express.json());

app.get("/", (req, res) => {
  res.send("🚀 Ingestion Service Running");
});

app.use(sensorRoutes);

export default app;