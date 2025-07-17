package llm

import (
	"github.com/google/generative-ai-go/genai"
)

var SubmitChoicesFunc = &genai.FunctionDeclaration{
	Name:        "submit_choices_and_story",
	Description: GenerateChoicesPrompt,
	Parameters: &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"narrative_text": {
				Type:        genai.TypeString,
				Description: "The narrative text for the next turn of the story.",
			},
			"narrative_summary": {
				Type:        genai.TypeString,
				Description: "A brief summary of the narrative so far, used for context in future turns.",
			},
			"choices": {
				Type:        genai.TypeArray,
				Description: "An array of three possible choices for the player.",
				Items: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"text": {Type: genai.TypeString, Description: "The text of the choice (e.g., 'Go into the cave')."},
					},
					Required: []string{"text"},
				},
			},
		},
		Required: []string{"narrative_text", "choices"},
	},
}
