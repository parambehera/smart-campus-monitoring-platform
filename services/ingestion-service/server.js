import app from "./src/app.js";
import { env } from "./src/config/env.js";
import { connectProducer } from "./src/kafka/producer.js";

async function startServer() {
  try {
    await connectProducer();

    app.listen(env.PORT, () => {
      console.log(
        `🚀 ${env.SERVICE_NAME} running on port ${env.PORT}`
      );
    });
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
}

startServer();