import { Router } from "express";
import { receiveSensorData } from "../controllers/sensor.controller.js";

const router = Router();

router.post("/sensor-data", receiveSensorData);

export default router;