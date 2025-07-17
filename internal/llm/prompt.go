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
const startAdventurePrompt = "Your task is to start a new adventure for the character named {{.Name}}. Describe the opening scene, introduce a mysterious hook, and leave the story open-ended." + "\n" + submitChoicesAndNarrativePrompt

const interactPrompt = `
	You are continuing the story for the character named {{.Name}}.
	The story so far is: "{{.NarrativeSummary}}"
	The player has just chosen to: "{{.Choice}}"
` + "\n" + submitChoicesAndNarrativePrompt

const submitChoicesAndNarrativePrompt = `
	RULE: You must respond by calling the 'submit_choices_and_story' function.

	ARGUMENTS:
	- 'narrative_text': The story text for the current turn.
	- 'updated_summary': A brief summary of the entire story so far, including this turn. If given a previous summary, append to it.
	- 'choices': An array of 3 distinct player choices.

	---
	**EXAMPLE**
	**TASK:** Start an adventure for Kaelen.
	**RESPONSE:**
	{
		"function_call": {
			"name": "submit_choices_and_story",
			"args": {
				"narrative_text": "You are Kaelen, in a rain-slicked Silverport alley. A note reads: 'Find the Crimson Canary'. What's your move?",
				"updated_summary": "Kaelen, a rogue in Silverport, has just received a mysterious note about the 'Crimson Canary'.",
				"choices": [
					{"text": "Go to the docks."},
					{"text": "Visit the tavern."},
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

func StartAdventurePrompt(characterName string) string {
	tmpl, _ := template.New("start").Parse(startAdventurePrompt)
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{"Name": characterName})
	return buf.String()
}

func InteractPrompt(characterName string, narrativeSummary string, choice string) string {
	tmpl, _ := template.New("interact").Parse(interactPrompt)
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]string{
		"Name":             characterName,
		"NarrativeSummary": narrativeSummary,
		"Choice":           choice,
	})
	return buf.String()
}

// Always include the gamemaster prompt at the start of every interaction
func Prompt(additionalPrompts []string) string {
	return GamemasterPrompt + "\n" + strings.Join(additionalPrompts, "\n")
}
