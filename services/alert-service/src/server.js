const express = require("express");
require("dotenv").config();

const healthRoute = require("./routes/health");

const app = express();

app.use(express.json());

const PORT = process.env.PORT || 5004;

app.get("/", (req, res) => {
  res.send("Alert Service Running");
});

app.use("/health", healthRoute);

app.listen(PORT, () => {
  console.log(`Alert Service running on port ${PORT}`);
});