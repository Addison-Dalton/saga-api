package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Addison-Dalton/saga-api/internal/shared"
	"github.com/google/generative-ai-go/genai"
)

type Service struct {
	genaiModel *genai.GenerativeModel
}

func NewService(genaiModel *genai.GenerativeModel) *Service {
	return &Service{
		genaiModel: genaiModel,
	}
}

func (s Service) GenerateStoryTurn(prompt string) (*shared.StoryTurn, error) {
	// 1. Configure the model with the necessary tools for this type of request
	modelWithTools := s.genaiModel
	modelWithTools.Tools = []*genai.Tool{
		{FunctionDeclarations: []*genai.FunctionDeclaration{SubmitChoicesFunc}},
	}

	// 2. Call the Gemini API
	resp, err := modelWithTools.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}

	// 3. Parse the response
	if resp != nil && len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if fc, ok := part.(genai.FunctionCall); ok && fc.Name == "submit_choices_and_story" {
			var turnData shared.StoryTurn // Use the struct defined in the game package
			argsBytes, _ := json.Marshal(fc.Args)
			if err := json.Unmarshal(argsBytes, &turnData); err == nil {
				return &turnData, nil // Success! Return the parsed data.
			}
		}
	}
	return nil, fmt.Errorf("LLM did not return a valid function call")
}
