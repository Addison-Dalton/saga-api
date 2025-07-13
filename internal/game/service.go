package game

import (
	"github.com/Addison-Dalton/saga-api/internal/storage"
	"github.com/google/generative-ai-go/genai"
)

type Service struct {
	DB            *storage.Database
	genaiModel    *genai.GenerativeModel
	activeSession *Session
}

func NewService(db *storage.Database, genaiModel *genai.GenerativeModel) *Service {
	return &Service{
		DB:            db,
		genaiModel:    genaiModel,
		activeSession: nil, // No active session at initialization
	}
}
