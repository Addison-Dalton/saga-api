package llm

import (
	"context"
	"log"

	"github.com/Addison-Dalton/saga-api/internal/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InitializeGenAIClient() (*genai.Client, error) {
	geminiAPIKey := config.Get("GEMINI_API_KEY")
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiAPIKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	return client, nil
}
