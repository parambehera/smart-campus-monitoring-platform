import { z } from "zod";

export const sensorSchema = z.object({
  deviceId: z.string().min(1),

  sensorType: z.enum([
    "Temperature",
    "Humidity",
    "Motion",
    "Smoke",
  ]),

  building: z.string().min(1),

  value: z.union([
    z.number(),
    z.boolean()
  ]),

  unit: z.string().optional(),

  battery: z.number().min(0).max(100),
});