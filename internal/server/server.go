package server

import (
	"github.com/Addison-Dalton/saga-api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type Server struct {
	router      *gin.Engine
	genaiClient *genai.GenerativeModel
	db          *storage.Database
}

func NewServer(genaiClient *genai.GenerativeModel, db *storage.Database) *Server {
	router := gin.Default()

	s := &Server{
		router:      router,
		genaiClient: genaiClient,
		db:          db,
	}

	api := s.router.Group("/api/v1")
	{
		// Test route for Gemini API
		api.POST("/test-prompt", s.testPromptHandler)
		// Character routes
		characters := api.Group("/characters")
		{
			characters.POST("/", s.CreateCharacterHandler)
			characters.GET("/", s.GetAllCharactersHandler)
			// TODO Get by ID, Update, Delete
		}
	}

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
