package game

import (
	"errors"
	"log"

	"github.com/Addison-Dalton/saga-api/internal/llm"
	"github.com/Addison-Dalton/saga-api/internal/shared"
	"github.com/Addison-Dalton/saga-api/internal/storage"
)

type Session struct {
	Character        storage.Character
	NarrativeSummary string
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

func (s *Service) CreateNewStory() (*shared.StoryTurn, error) {
	// validate that there is an active session
	if s.activeSession == nil {
		return nil, ErrNoActiveSession
	}

	// Generate a new story prompt using the character's details
	newStoryPrompt := llm.StartAdventurePrompt(s.activeSession.Character.Name)

	storyTurn, err := s.llm.GenerateStoryTurn(newStoryPrompt)

	if err != nil {
		log.Printf("Error generating story turn: %v", err)
		return nil, err
	}

	// Update the session's narrative summary with the new story turn
	s.activeSession.NarrativeSummary = storyTurn.NarrativeSummary

	return storyTurn, nil
}

func (s *Service) Interact(choice string) (*shared.StoryTurn, error) {
	// validate that there is an active session
	if s.activeSession == nil {
		return nil, ErrNoActiveSession
	}

	// Prepare the interaction prompt
	interactionPrompt := llm.InteractPrompt(s.activeSession.Character.Name, s.activeSession.NarrativeSummary, choice)

	storyTurn, err := s.llm.GenerateStoryTurn(interactionPrompt)

	if err != nil {
		log.Printf("Error generating story turn: %v", err)
		return nil, err
	}

	// Update the session's narrative summary with the new story turn
	s.activeSession.NarrativeSummary = storyTurn.NarrativeSummary

	return storyTurn, nil
}
