const express = require("express");
const router = express.Router();

const { sendEmail } = require("../services/emailService");

router.get("/", async (req, res) => {
  try {
    await sendEmail(
      "dasashutosh606@gmail.com",
      "Test Email from Alert Service",
      "Hello! Your Alert Service email is working successfully."
    );

    res.json({
      message: "Email sent successfully"
    });
  } catch (error) {
    console.error(error);

    res.status(500).json({
      message: "Failed to send email"
    });
  }
});

module.exports = router;