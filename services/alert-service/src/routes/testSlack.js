const express = require("express");
const router = express.Router();

const {
  sendSlackMessage,
} = require("../services/slackService");

router.get("/", async (req, res) => {
  try {
    await sendSlackMessage(
      "🚨 Test Alert from Smart Campus Alert Service!"
    );

    res.json({
      message: "Slack message sent successfully",
    });
  } catch (error) {
    res.status(500).json({
      message: "Failed to send Slack message",
    });
  }
});

module.exports = router;