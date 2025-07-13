package llm

import (
	"bytes"
	"strings"
	"text/template"
)

// LLM prompts

// GamemasterPrompt is the overarching prompt for the LLM to act as a fantasy RPG gamemaster.
const GamemasterPrompt = "You are a fantasy RPG gamemaster for a text-based game. Narrate immersive events and play all NPCs. Never break character."

// StartAdventurePrompt is used to generate a compelling starting scenario for a new adventure.
// The placeholder {{.Name}} will be replaced with the character's name.
const StartAdventurePrompt = "Your task is to start a new adventure for the character named {{.Name}}. Describe the opening scene, introduce a mysterious hook, and leave the story open-ended." + "\n" + SubmitChoicesAndNarrativePrompt

const SubmitChoicesAndNarrativePrompt = `
You MUST call the 'submit_choices_and_story' function to provide your response.

---
**EXAMPLE**
**TASK:** Start an adventure for Kaelen.
**RESPONSE (Function Call):**
{
  "function_call": {
    "name": "submit_choices_and_story",
    "args": {
      "narrative_text": "Kaelen, you're in a rain-slicked alley in Silverport. A note under your door reads: 'Find the Crimson Canary'. What do you do?",
      "choices": [
        {"text": "Go to the docks."},
        {"text": "Visit the local tavern."},
        {"text": "Examine the note."}
      ]
    }
  }
}
---
`

// GenerateChoicesPrompt instructs the LLM on how to generate choices for the player.
const GenerateChoicesPrompt = `
Submit the story text and the 3 choices for the player.
`

func GetStartAdventurePrompt(characterName string) string {
	tmpl, _ := template.New("start").Parse(StartAdventurePrompt)
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"Name": characterName})
	return buf.String()
}

// Always include the gamemaster prompt at the start of every interaction
func Prompt(additionalPrompts []string) string {
	return GamemasterPrompt + "\n" + strings.Join(additionalPrompts, "\n")
}
