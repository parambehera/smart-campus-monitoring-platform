const { IncomingWebhook } = require("@slack/webhook");
require("dotenv").config();
console.log(process.env.SLACK_WEBHOOK_URL);
const webhook = new IncomingWebhook(
  process.env.SLACK_WEBHOOK_URL
);

async function sendSlackMessage(message) {
  try {
    await webhook.send({
      text: message,
    });

    console.log("Slack message sent successfully");
  } catch (error) {
    console.error("Slack message failed:", error);
  }
}

module.exports = { sendSlackMessage };