const kafka = require("../config/kafka");
const { sendEmail } = require("../services/emailService");
const { sendSlackMessage } = require("../services/slackService");

const consumer = kafka.consumer({
  groupId: "alert-group",
});

async function startConsumer() {
  await consumer.connect();

  console.log("Kafka Consumer Connected");

  await consumer.subscribe({
    topic: "alerts",
    fromBeginning: true,
  });

  await consumer.run({
    eachMessage: async ({ message }) => {
      try {
        const alert = JSON.parse(
          message.value.toString()
        );

        console.log("Alert Received:", alert);

        await sendEmail(
          process.env.ADMIN_EMAIL,
          "Campus Alert",
          alert.message
        );

        await sendSlackMessage(
          `🚨 ${alert.message}`
        );
      } catch (err) {
        console.error("Error:", err);
      }
    },
  });
}

module.exports = startConsumer;