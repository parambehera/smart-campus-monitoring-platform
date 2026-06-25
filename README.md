# smart-campus-monitoring-platform

A scalable **microservices-based IoT monitoring platform** developed as part of the **TCS Xcelerate Industry-Aligned Capstone Program**. The platform collects real-time sensor telemetry from multiple campus buildings to monitor electricity consumption, water utilization, and indoor air quality. It processes time-series data, performs analytics, detects anomalies, generates alerts, and provides live dashboards using an event-driven architecture powered by MQTT and Apache Kafka.

---

# 🚀 Technology Stack

| Category             | Technology          |
| -------------------- | ------------------- |
| Frontend             | React, Tailwind CSS |
| Backend              | Go, Python, Node.js |
| IoT Communication    | MQTT (Mosquitto)    |
| Event Streaming      | Apache Kafka        |
| Time-Series Database | TimescaleDB         |
| Cache                | Redis               |
| Background Jobs      | Celery              |
| API Gateway          | Kong                |
| Containerization     | Docker              |
| Orchestration        | Kubernetes          |
| Package Manager      | Helm                |
| Monitoring           | Prometheus          |
| Visualization        | Grafana             |
| Version Control      | Git & GitHub        |

---

# 🏗️ Project Structure

```text
smart-campus-monitoring-platform/
│
├── assets/                     # Images, logos & diagrams
├── docs/                       # Project documentation
├── frontend/
│   └── react-dashboard/        # React frontend
│
├── services/
│   ├── device-service/
│   ├── ingestion-service/
│   ├── analytics-service/
│   ├── alert-service/
│   └── dashboard-service/
│
├── simulator/
│   └── sensor-simulator/
│
├── infrastructure/
│   ├── mqtt/
│   ├── kafka/
│   ├── redis/
│   ├── timescaledb/
│   ├── kong/
│   ├── prometheus/
│   ├── grafana/
│   └── kubernetes/
│
├── helm/
├── docker-compose.yml
├── .env.example
├── README.md
├── LICENSE
└── .gitignore
```

---

# ⚙️ Microservices

| Service           | Technology                  | Responsibility                               |
| ----------------- | --------------------------- | -------------------------------------------- |
| Device Service    | Go + Redis                  | Sensor registration and metadata management  |
| Ingestion Service | Python + MQTT + TimescaleDB | Receives sensor data and stores telemetry    |
| Analytics Service | Python + Celery             | Hourly/Daily analytics and anomaly detection |
| Alert Service     | Node.js                     | Sends Email and Slack alerts                 |
| Dashboard Service | Node.js + WebSocket         | REST APIs and real-time dashboard            |
| Frontend          | React                       | Live monitoring dashboard                    |

---

# 🔄 System Flow

```text
IoT Sensors
     │
     ▼
 MQTT Broker (Mosquitto)
     │
     ▼
 Ingestion Service
     │
     ├────────────► TimescaleDB
     │
     ▼
 Apache Kafka
     │
 ┌───┼───────────────┐
 │   │               │
 ▼   ▼               ▼
Analytics     Alert Service    Dashboard Service
                                  │
                                  ▼
                           React Dashboard
```

---

# 📌 Key Features

* Microservices Architecture
* MQTT-based IoT Communication
* Apache Kafka Event Streaming
* Time-Series Data Storage
* Real-time WebSocket Dashboard
* Automated Alert Generation
* Asynchronous Analytics Processing
* Kubernetes Deployment
* Prometheus & Grafana Monitoring
* Scalable and Event-Driven Design

---




# 🤝 Git Workflow & Contribution Guide

This project follows a **Git Flow** branching strategy to ensure smooth collaboration and avoid merge conflicts. Every team member must follow the workflow below.

---

# Repository Branches

| Branch      | Purpose                                                                            |
| ----------- | ---------------------------------------------------------------------------------- |
| `main`      | Stable and production-ready code. Only the Team Leader can merge into this branch. |
| `develop`   | Integration branch where completed features from all developers are merged.        |
| `feature/*` | Individual feature branches created by each developer for their assigned tasks.    |

---

# Step 1: Clone the Repository

Clone the repository to your local machine.

```bash
git clone https://github.com/parambehera/smart-campus-monitoring-platform.git
```

Move into the project directory.

```bash
cd smart-campus-monitoring-platform
```

---

# Step 2: Switch to the Development Branch

After cloning, switch to the `develop` branch.

```bash
git checkout develop
```

Always start your work from the latest version of the `develop` branch.

---

# Step 3: Create Your Own Feature Branch

Create a separate branch for your assigned module.

Examples:

```bash
git checkout -b feature/device-service
```

```bash
git checkout -b feature/ingestion-service
```

```bash
git checkout -b feature/analytics-service
```

```bash
git checkout -b feature/alert-service
```

```bash
git checkout -b feature/dashboard-service
```

Each developer should work only in their own feature branch.

---

# Step 4: Implement Your Feature

Develop only the module assigned to you.

Examples:

* Device Service
* Ingestion Service
* Analytics Service
* Alert Service
* Dashboard Service

Do not modify another developer's module unless discussed with the team.

---

# Step 5: Commit Your Changes

Stage all files.

```bash
git add .
```

Commit your changes with a meaningful message.

Example:

```bash
git commit -m "Implement device registration API"
```

Always use descriptive commit messages.

---

# Step 6: Push Your Feature Branch

Push your feature branch to GitHub.

Example:

```bash
git push origin feature/device-service
```

Do **NOT** push directly to:

* `main`
* `develop`

Only push to your own feature branch.

---

# Step 7: Create a Pull Request

After pushing your feature branch:

1. Open the GitHub repository.
2. Click **Compare & Pull Request**.
3. Set:

```
Base Branch  : develop
Compare Branch : feature/<your-feature>
```

4. Submit the Pull Request.

---

# Step 8: Code Review

The Team Leader will:

* Review the code.
* Verify folder structure.
* Check coding standards.
* Suggest changes if required.

If everything is correct, the Pull Request will be merged into the `develop` branch.

---

# Step 9: Update Your Local Repository

Before starting new work, always update your local `develop` branch.

```bash
git checkout develop
```

```bash
git pull origin develop
```

Then switch back to your feature branch if needed.

---

# Step 10: Continue Development

Continue implementing your assigned feature.

Repeat the cycle:

1. Code
2. Commit
3. Push
4. Create Pull Request
5. Wait for Review
6. Merge into `develop`

---

# Final Release Process

Once every feature has been completed and tested:

```
feature branches
        ↓
     develop
        ↓
 Integration Testing
        ↓
       main
```

Only the Team Leader is responsible for merging the `develop` branch into the `main` branch after the project is stable.

---

# Team Rules

✅ Always create a new feature branch from `develop`.

✅ Keep commit messages meaningful.

✅ Pull the latest `develop` branch before starting new work.

✅ Create a Pull Request for every completed feature.

❌ Never push directly to `main`.

❌ Never push directly to `develop`.

❌ Never modify another developer's module without discussion.

---

# Example Branch Names

```
feature/device-service
feature/ingestion-service
feature/analytics-service
feature/alert-service
feature/dashboard-service
feature/frontend
feature/documentation
feature/kubernetes
```

Following this workflow ensures a clean Git history, simplifies collaboration, and minimizes merge conflicts throughout the project.
