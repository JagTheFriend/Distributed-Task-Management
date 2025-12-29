# Distributed Task Queue (Golang)

A **production-inspired distributed task queue** built in **Golang** using **Redis**, designed to demonstrate real-world backend engineering concepts such as concurrency, fault tolerance, retries, and crash recovery.

This project is **purely backend-focused** and intentionally avoids frontend concerns to highlight system design and backend architecture.

---

## 🚀 Features

* Asynchronous task execution
* Distributed, stateless workers
* At-least-once delivery guarantees
* Visibility timeout & crash recovery
* Retry mechanism with exponential backoff
* Dead-letter queue for failed tasks
* Clean, modular Go project structure
* Redis-backed coordination

---

## 🧠 Why This Project Exists

In real systems, long-running or unreliable tasks should not block user requests.
A distributed task queue enables:

* Decoupling producers from consumers
* Horizontal scalability
* Fault tolerance
* Better system reliability

This architecture is inspired by systems like **Celery, Sidekiq, AWS SQS, and Google Cloud Tasks**.

---

## 🏗 Architecture Overview

```
Client
  |
  v
API Server (Producer)
  |
  v
Redis
 ├── Pending Queue
 ├── Running Set (visibility timeout)
 ├── Retry Set (delayed retries)
 └── Dead Letter Queue
  |
  v
Worker Pool (Consumers)
```

### Task Lifecycle

```
PENDING → RUNNING → DONE
              ↓
           FAILED → RETRY → DEAD
```

---

## 📁 Project Structure

```
task-queue/
├── cmd/
│   ├── api/        # API server (task producer)
│   └── worker/     # Worker service (task consumer)
├── internal/
│   ├── api/        # HTTP handlers
│   ├── config/     # Redis configuration
│   ├── queue/      # Redis queue abstraction
│   ├── task/       # Task model, repository, service
│   └── worker/     # Task execution & reaper
├── go.mod
└── README.md
```

This structure follows **standard Go project conventions** and clean separation of concerns.

---

## 🧩 Core Components

### API Server

* Accepts task submissions via HTTP
* Persists task metadata
* Enqueues tasks into Redis

### Redis

Used as a coordination layer:

* `task:pending` → ready-to-run tasks
* `task:running` → tasks currently being processed
* `task:retry` → delayed retries
* `task:dead` → permanently failed tasks

### Workers

* Pull tasks from Redis
* Execute tasks concurrently
* Handle retries and failures
* Are stateless and horizontally scalable

### Reaper

* Detects tasks stuck due to worker crashes
* Re-queues expired tasks safely

---

## ⚙️ Tech Stack

* **Language:** Go (1.20+)
* **Queue & Coordination:** Redis
* **Concurrency:** Goroutines
* **API:** net/http
* **Serialization:** JSON

---

## ▶️ Getting Started

### Prerequisites

* Go 1.20+
* Redis running locally on `localhost:6379`

---

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/yourusername/task-queue.git
cd task-queue
```

---

### 2️⃣ Start the API Server

```bash
go run cmd/api/main.go
```

Server runs on:

```
http://localhost:8080
```

---

### 3️⃣ Start a Worker

```bash
go run cmd/worker/main.go
```

You can start **multiple workers** in separate terminals to simulate horizontal scaling.

---

### 4️⃣ Submit a Task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"type":"email","payload":"hello world"}'
```

---

## 🔁 Reliability Guarantees

* **At-least-once execution**
* Tasks may be retried on failure
* Workers are safe to crash and restart
* Duplicate execution is possible → tasks must be idempotent

This mirrors real-world distributed systems behavior.

---

## 📈 Scalability

* Add more worker processes to increase throughput
* Redis acts as a central coordination point
* Workers remain stateless
* Supports backpressure and retry delays

---

## 🧪 Failure Handling

* Failed tasks are retried with exponential backoff
* Tasks exceeding retry limits are moved to a dead-letter queue
* Expired running tasks are recovered automatically

---

## 🎯 What This Project Demonstrates

* Distributed systems fundamentals
* Concurrency and worker pools in Go
* Redis data structures (Lists, ZSETs)
* Fault tolerance and crash recovery
* Clean architecture and package design
* Production-style backend thinking

---

## 🔮 Future Improvements

* Priority queues
* gRPC API
* Metrics & Prometheus integration
* Idempotency keys
* Persistent database backing
* Web dashboard
* Kubernetes deployment

---
