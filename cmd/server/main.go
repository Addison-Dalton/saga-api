package main

import (
	"context"
	"log"

	"github.com/Addison-Dalton/saga-api/internal/config"
	"github.com/Addison-Dalton/saga-api/internal/server"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Prompt struct {
	Message string `json:"message"`
}

func main() {
	// environment variable loading
	config.Load()
	geminiAPIKey := config.Get("GEMINI_API_KEY")

	// gemini client initialization
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiAPIKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	geminiModel := client.GenerativeModel("gemini-1.5-flash-latest")

	srv := server.NewServer(geminiModel)
	log.Printf("Starting server on %s", ":8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
