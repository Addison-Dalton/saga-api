package server

import (
	"github.com/Addison-Dalton/saga-api/internal/game"
	"github.com/Addison-Dalton/saga-api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type Server struct {
	router      *gin.Engine
	genaiClient *genai.GenerativeModel
	db          *storage.Database
	gameService *game.Service
}

func NewServer(genaiClient *genai.GenerativeModel, db *storage.Database, gameService *game.Service) *Server {
	router := gin.Default()

	s := &Server{
		router:      router,
		genaiClient: genaiClient,
		db:          db,
		gameService: gameService,
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
			characters.GET("/:id", s.GetCharacterByIDHandler)
			// TODO Get by Update, Delete
		}
		// game routes
		gameRoutes := api.Group("/game")
		{
			gameRoutes.POST("/start", s.SessionStartHandler)
			// gameRoutes.POST("/interact", s.GamePromptHandler)
		}
	}

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
