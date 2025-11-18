package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	JWTSecret  string
	JWTExpiry  time.Duration
	JWTIssuer  string
	ServerPort string
	DB         *gorm.DB
}

func Load() *Config {
	_ = godotenv.Load(".env")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	expStr := os.Getenv("JWT_EXP_MINUTES")
	if expStr == "" {
		expStr = "15"
	}
	expMin, err := strconv.Atoi(expStr)
	if err != nil {
		log.Fatalf("invalid JWT_EXP_MINUTES: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := initDB()

	return &Config{
		JWTSecret:  jwtSecret,
		JWTExpiry:  time.Duration(expMin) * time.Minute,
		JWTIssuer:  os.Getenv("JWT_ISSUER"),
		ServerPort: ":" + port,
		DB:         db,
	}
}

func initDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")
	ssl := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		host, port, user, pass, name, ssl, schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL.")
	return db
}
