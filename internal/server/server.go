package server

import (
	"github.com/Addison-Dalton/saga-api/internal/game"
	"github.com/Addison-Dalton/saga-api/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	db          *storage.Database
	gameService *game.Service
}

func NewServer(db *storage.Database, gameService *game.Service) *Server {
	router := gin.Default()

	s := &Server{
		router:      router,
		db:          db,
		gameService: gameService,
	}

	api := s.router.Group("/api/v1")
	{
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
			gameRoutes.POST("/interact", s.SessionInteractHandler)
		}
	}

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
