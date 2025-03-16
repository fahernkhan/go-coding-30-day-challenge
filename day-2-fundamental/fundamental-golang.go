// # File: main.go
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var ctx = context.Background()

// UserService handles user-related operations
type UserService struct {
	DB    *sql.DB       // Database connection
	Cache *redis.Client // Redis cache client
}

// NewUserService creates a new UserService instance with Dependency Injection
func NewUserService(db *sql.DB, cache *redis.Client) *UserService {
	return &UserService{DB: db, Cache: cache}
}

var userService *UserService

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database and Redis cache
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	cache := initRedis()
	defer cache.Close()

	// Initialize user service with DI
	userService = NewUserService(db, cache)

	// Setup Gin Router
	r := gin.Default()
	r.Use(RequestLoggerMiddleware()) // Middleware for logging requests
	r.Use(CORSMiddleware())          // Middleware for handling CORS

	// Public routes
	r.POST("/register", userService.registerUser) // Endpoint for user registration
	r.POST("/login", userService.loginUser)       // Endpoint for user login
	r.GET("/ping", func(c *gin.Context) {         // Health check endpoint
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Protected routes requiring authentication
	authRoutes := r.Group("/api")
	authRoutes.Use(AuthMiddleware())
	authRoutes.GET("/profile", userService.getUserProfile)  // Fetch user profile
	authRoutes.GET("/users", userService.getAllUsers)       // Fetch all users
	authRoutes.DELETE("/users/:id", userService.deleteUser) // Delete a user by ID
	authRoutes.POST("/logout", userService.logoutUser)      // User logout

	// Run server on specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}

// Initialize PostgreSQL database connection
func initDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Initialize Redis cache connection
func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
}

// Middleware to log incoming requests
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("%s %s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr, c.Writer.Status(), time.Since(start))
	}
}

// Middleware to handle CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// Logout user and invalidate session
func (s *UserService) logoutUser(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	s.Cache.Set(ctx, fmt.Sprintf("token_%s", username), "", time.Minute*5) // Invalidate token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Middleware for JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims := &jwt.MapClaims{}
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", (*claims)["username"])
		c.Next()
	}
}

// Error handling for user not found
var ErrUserNotFound = errors.New("user not found")

// Fetch stored password hash for a user
func (s *UserService) getUserByUsername(username string) (string, error) {
	var storedPassword string
	err := s.DB.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUserNotFound
		}
		return "", err
	}
	return storedPassword, nil
}

// Login user and generate JWT token
func (s *UserService) loginUser(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedPassword, err := s.getUserByUsername(user.Username)
	if err != nil {
		if err == ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
