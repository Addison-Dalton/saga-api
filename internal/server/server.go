package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type Server struct {
	router      *gin.Engine
	genaiClient *genai.GenerativeModel
}

func NewServer(genaiClient *genai.GenerativeModel) *Server {
	router := gin.Default()

	s := &Server{
		router:      router,
		genaiClient: genaiClient,
	}

	s.router.POST("/test-prompt", s.testPromptHandler)

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
