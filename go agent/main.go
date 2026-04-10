package main

import (
	"go-agent/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, make sure HUGGINGFACE_API_KEY is set in env")
	}
}
func main() {
	apiKey := os.Getenv("HUGGINGFACE_API_KEY")
	if apiKey == "" {
		log.Fatal("HUGGINGFACE_API_KEY is not set!")
	}
	log.Println("HF API key loaded ✅")
	r := gin.Default()
	routes.GetVacationRouter(r)
	r.Run()
}
