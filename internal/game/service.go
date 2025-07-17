package game

import (
	"github.com/Addison-Dalton/saga-api/internal/llm"
	"github.com/Addison-Dalton/saga-api/internal/storage"
)

type Service struct {
	DB            *storage.Database
	llm           *llm.Service
	activeSession *Session
}

func NewService(db *storage.Database, llm *llm.Service) *Service {
	return &Service{
		DB:            db,
		llm:           llm,
		activeSession: nil, // No active session at initialization
	}
}
