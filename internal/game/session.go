package game

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Addison-Dalton/saga-api/internal/llm"
	"github.com/Addison-Dalton/saga-api/internal/storage"

	"github.com/google/generative-ai-go/genai"
)

type Session struct {
	Character        storage.Character
	NarrativeSummary string
}

type Choice struct {
	Text string `json:"text"`
}

type StoryTurn struct {
	NarrativeText string   `json:"narrative_text"`
	Choices       []Choice `json:"choices"`
}

var ErrNoActiveSession = errors.New("no active game session")

// NewSession initializes a new game session with the given character ID.
func (s *Service) NewSession(characterID uint) error {
	// Reset the session if it already exists
	if s.activeSession != nil {
		s.activeSession = nil
	}

	// Fetch the character from the database
	character, err := s.DB.GetCharacterByID(characterID)
	if err != nil {
		return err
	}

	session := &Session{
		Character:        *character,
		NarrativeSummary: "",
	}
	s.activeSession = session

	return nil
}

func (s *Service) CreateNewStory() (*StoryTurn, error) {
	// validate that there is an active session
	if s.activeSession == nil {
		return nil, ErrNoActiveSession
	}

	// Generate a new story prompt using the character's details
	newStoryPrompt := llm.GetStartAdventurePrompt(s.activeSession.Character.Name)

	modelWithTools := s.genaiModel
	modelWithTools.Tools = []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{llm.SubmitChoicesFunc},
		},
	}

	resp, err := modelWithTools.GenerateContent(
		context.Background(),
		genai.Text(llm.Prompt([]string{newStoryPrompt})),
	)

	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}

	if resp != nil && len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		part := resp.Candidates[0].Content.Parts[0]

		if fc, ok := part.(genai.FunctionCall); ok && fc.Name == "submit_choices_and_story" {

			var turnData StoryTurn

			argsBytes, _ := json.Marshal(fc.Args)
			if err := json.Unmarshal(argsBytes, &turnData); err == nil {
				return &StoryTurn{
					NarrativeText: turnData.NarrativeText,
					Choices:       turnData.Choices,
				}, nil
			} else {
				log.Printf("Failed to unmarshal function call args: %v", err)
			}
		}
	}

	return nil, fmt.Errorf("LLM did not return a valid function call")
}
