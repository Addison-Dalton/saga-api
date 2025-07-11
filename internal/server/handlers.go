package server

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Addison-Dalton/saga-api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
)

type Prompt struct {
	Message string `json:"message"`
}

func (s *Server) testPromptHandler(c *gin.Context) {
	var req Prompt

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	log.Printf("Received prompt: %s", req.Message)

	// Use gemini client to process the prompt
	resp, err := s.genaiClient.GenerateContent(context.Background(), genai.Text(req.Message))
	if err != nil {
		log.Printf("Gemini API error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate content"})
		return
	}

	// just grabbing the first response for test purposes
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			c.JSON(http.StatusOK, gin.H{"response": string(text)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"response": "Recieved prompt successfully, but no valid response from Gemini"})
}

// Character handlers
func (s *Server) CreateCharacterHandler(c *gin.Context) {
	var character storage.Character
	if err := c.BindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	if err := s.db.CreateCharacter(&character); err != nil {
		log.Printf("Error creating character: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create character"})
		return
	}

	c.JSON(http.StatusCreated, character)
}

func (s *Server) GetAllCharactersHandler(c *gin.Context) {
	characters, err := s.db.GetAllCharacters()

	if err != nil {
		log.Printf("Error retrieving characters: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve characters"})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (s *Server) GetCharacterByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	// Step 2: Convert the 'id' to uint
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert to uint32
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	character, dbErr := s.db.GetCharacterByID(uint(id))

	if dbErr != nil {
		log.Printf("Error retrieving character with ID %d: %v", id, dbErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve character"})
		return
	}

	c.JSON(http.StatusOK, character)
}

// TODO UpdateCharacter, DeleteCharacter handlers
