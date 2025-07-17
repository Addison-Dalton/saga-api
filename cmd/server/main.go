package main

import (
	"log"

	"github.com/Addison-Dalton/saga-api/internal/config"
	"github.com/Addison-Dalton/saga-api/internal/game"
	"github.com/Addison-Dalton/saga-api/internal/llm"
	"github.com/Addison-Dalton/saga-api/internal/server"
	"github.com/Addison-Dalton/saga-api/internal/storage"
)

type Prompt struct {
	Message string `json:"message"`
}

func main() {
	// environment variable loading
	config.Load()
	modelName := config.Get("GEMINI_MODEL")
	dbConnectionString := config.Get("DATABASE_URL")

	// database initialization
	db, _ := storage.NewConnection(dbConnectionString)
	storage.AutoMigrate(db.DB)

	// gemini client initialization
	client, err := llm.InitializeGenAIClient()
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	// llm service initialization
	llmService := llm.NewService(model)

	// game service initialization
	gameService := game.NewService(db, llmService)

	srv := server.NewServer(db, gameService)
	log.Printf("Starting server on %s", ":8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
