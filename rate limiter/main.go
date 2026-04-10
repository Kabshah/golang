package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	APIKey    string    `gorm:"uniqueIndex;size:255" json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
}

type RateLimit struct {
	Key         string    `gorm:"primaryKey;size:255" json:"key"`
	Count       int       `json:"count"`
	WindowStart time.Time `gorm:"index" json:"window_start"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	db.AutoMigrate(&User{}, &RateLimit{})
	return db
}

func generateAPIKey() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)
	return "API-" + hex.EncodeToString(bytes)
}

// main rate limit logic

func RateLimiter(db *gorm.DB, limit int, window time.Duration, exceedStatus int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			
		}
	}
}
