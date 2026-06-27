# Alert Service

## Overview

The Alert Service is responsible for consuming alert messages from Apache Kafka and sending notifications to administrators through Email and Slack.

## Features

* Health Check API
* Email Notification Service
* Slack Notification Service
* Kafka Consumer Integration
* Docker Support

---

# Architecture

Kafka Topic (`alerts`)
→ Alert Service
→ Email Notifications
→ Slack Notifications

---

# Technology Stack

* Node.js
* Express.js
* Apache Kafka (KafkaJS)
* Nodemailer
* Slack Incoming Webhooks
* Docker

---

# Project Structure

```text
alert-service/
│
├── src/
│   ├── config/
│   │   └── kafka.js
│   │
│   ├── consumers/
│   │   └── alertConsumer.js
│   │
│   ├── routes/
│   │   ├── health.js
│   │   ├── testEmail.js
│   │   └── testSlack.js
│   │
│   ├── services/
│   │   ├── emailService.js
│   │   └── slackService.js
│   │
│   └── server.js
│
├── Dockerfile
├── package.json
├── .env
└── README.md
```

---

# Installation

## Clone Repository

```bash
git clone <repository-url>
cd services/alert-service
```

## Install Dependencies

```bash
npm install
```

---

# Environment Variables

Create a `.env` file:

```env
PORT=5004
EMAIL_USER=your-email@gmail.com
EMAIL_PASS=your-app-password
ADMIN_EMAIL=admin@gmail.com
SLACK_WEBHOOK_URL=your-slack-webhook-url
```

---

# Running the Service

## Development Mode

```bash
npm run dev
```

## Production Mode

```bash
npm start
```

---

# API Endpoints

## Health Check

```http
GET /health
```

Response:

```json
{
  "status": "UP",
  "service": "alert-service",
  "timestamp": "2026-06-27T07:22:34.470Z"
}
```

---

## Test Email

```http
GET /test-email
```

Sends a test email to the configured recipient.

---

## Test Slack Notification

```http
GET /test-slack
```

Sends a test message to the configured Slack channel.

---

# Kafka Configuration

Broker:

```text
localhost:9092
```

Topic:

```text
alerts
```

Example Message:

```json
{
  "message": "Temperature exceeded threshold in Lab A"
}
```

---

# Docker Support

## Build Image

```bash
docker build -t alert-service .
```

## Run Container

```bash
docker run -p 5004:5004 --env-file .env alert-service
```

---

# Notification Flow

```text
Kafka Topic (alerts)
          │
          ▼
     Alert Service
          │
    ┌─────┴─────┐
    ▼           ▼
 Email       Slack
```

---

# Future Enhancements

* SMS Notifications
* Push Notifications
* Alert Retry Mechanism
* Alert History Storage
* Docker Compose Integration

---

# Contributors

* Ashutosh Das – Alert Service Development
* Smart Campus Monitoring Platform Team
