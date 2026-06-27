const express = require("express");
require("dotenv").config();

const healthRoute = require("./routes/health");
const testEmailRoute = require("./routes/testEmail");
const testSlackRoute = require("./routes/testSlack");

const app = express();

app.use(express.json());

const PORT = process.env.PORT || 5004;

app.get("/", (req, res) => {
  res.send("Alert Service Running");
});

app.use("/health", healthRoute);
app.use("/test-email", testEmailRoute);
app.use("/test-slack", testSlackRoute);

app.listen(PORT, () => {
  console.log(`Alert Service running on port ${PORT}`);
});