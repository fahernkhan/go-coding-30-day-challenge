# Golang Backend Development Fundamentals

## 📌 1. Struktur dan Arsitektur Backend di Golang
- **Layered Architecture** (MVC, Clean Architecture, Hexagonal, DDD)
- **Structuring Packages & Modules** (Best practice untuk modularisasi)
- **Error Handling** (`errors.Wrap`, `fmt.Errorf` dengan `%w`)
- **Logging** dengan `logrus` atau `zap`

## 📌 2. Routing dan Middleware
- **Routing** menggunakan `gorilla/mux` atau `gin-gonic/gin`
- **Middleware** (CORS, Rate Limiting, Request Logging, Authentication)
- **Custom Middleware** untuk kebutuhan spesifik
- **Context Management** (`context.Context`) untuk timeout dan pembatalan request

## 📌 3. Database Management
- **ORM vs SQL Native** (`gorm` vs `database/sql` + `sqlx`)
- **Migrations** (`golang-migrate`, `gorm` auto migrations)
- **Transactions Handling** (`tx.Begin()`, `tx.Commit()`, `tx.Rollback()`)
- **Connection Pooling** (`sql.DB.SetMaxOpenConns()`, `SetMaxIdleConns()`)
- **Indexing dan Query Optimization**
- **Handling Data Consistency dan Concurrency** (`SELECT FOR UPDATE`, `row-level locking`)

## 📌 4. Authentication & Authorization
- **JWT** (`github.com/golang-jwt/jwt`)
- **OAuth 2.0** (`golang.org/x/oauth2`)
- **Session-based Authentication** (Redis/Memcached)
- **RBAC (Role-Based Access Control) dan Claims**

## 📌 5. Caching dan Optimasi
- **Memanfaatkan `Redis` atau `Memcached`**
- **Caching strategi** (Lazy Loading, Write-Through, TTL Expiration)
- **Cache Invalidation** (Least Recently Used - LRU)
- **Pagination dengan `LIMIT OFFSET` dan Cursor Pagination**

## 📌 6. Asynchronous Processing & Job Queue
- **Background Job** (`go-routines`, `worker-pool`, `sync.WaitGroup`)
- **Queue Systems** (`RabbitMQ`, `Kafka`, `NATS`)
- **Task Scheduling** (`cron jobs`, `github.com/robfig/cron/v3`)

## 📌 7. API Development Best Practices
- **RESTful API Principles** (HTTP Methods, Status Codes, Idempotency)
- **gRPC** untuk komunikasi antar layanan (`google.golang.org/grpc`)
- **API Documentation** (`Swagger`, `Redoc`)
- **GraphQL dengan `gqlgen`**

## 📌 8. Microservices & Distributed Systems
- **Service Discovery** (`Consul`, `etcd`, `Zookeeper`)
- **gRPC dan Protobuf** untuk komunikasi antar service
- **Event-Driven Architecture** (`Kafka`, `NATS`, `RabbitMQ`)
- **Circuit Breaker** (`github.com/sony/gobreaker`)
- **Rate Limiting** (`golang.org/x/time/rate`)

## 📌 9. Observability (Monitoring, Logging, Tracing)
- **Structured Logging** dengan `logrus`, `zap`
- **Metrics Monitoring** (`Prometheus`, `Grafana`)
- **Distributed Tracing** (`Jaeger`, `OpenTelemetry`)
- **Health Check** (`github.com/heptiolabs/healthcheck`)

## 📌 10. Deployment & CI/CD
- **Dockerizing Golang Backend** (`Dockerfile`, multi-stage builds)
- **Kubernetes Deployment** (`Helm`, `Kustomize`)
- **CI/CD** (`GitHub Actions`, `GitLab CI`, `Jenkins`, `ArgoCD`)
- **Feature Flags** (`Unleash`, `LaunchDarkly`)
- **Configuration Management** (`Viper`, `.env`, `Secrets Management`)

## 📌 11. Security Best Practices
- **Input Validation** (`github.com/go-playground/validator/v10`)
- **SQL Injection Prevention** (`sqlx.NamedQuery`, parameterized queries)
- **XSS, CSRF, dan Security Headers** (`gin-secure`, `helmet`)
- **Secrets Management** (`AWS Secrets Manager`, `Vault`)
- **Secure File Uploads** dan **Content-Type Validation**

---

📂 golang-backend/
│   ├── 📂 config/        # Konfigurasi aplikasi (DB, Redis, dll.)
│   ├── 📂 controllers/   # Handler untuk routing API
│   ├── 📂 middlewares/   # Middleware (Auth, Logging, dll.)
│   ├── 📂 models/        # Model data (ORM / struct)
│   ├── 📂 repositories/  # Query ke database
│   ├── 📂 services/      # Business logic / service layer
│   ├── 📂 routes/        # Routing API
│   ├── 📂 utils/         # Helper functions (JWT, hashing, dll.)
│   ├── main.go          # Entry point aplikasi
│   ├── go.mod           # Dependency management
│   ├── go.sum           # Checksum dependencies
│   ├── .env             # Konfigurasi environment
│   ├── Dockerfile       # Dockerfile untuk deployment
│   ├── README.md        # Dokumentasi proyek

# File: main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db := config.InitDB()
	defer db.Close()

	// Setup Gin Router
	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}

# File: config/database.go
package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

# File: routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"golang-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}

# File: controllers/ping.go
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

# File: Dockerfile
FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080

