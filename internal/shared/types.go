package shared

type Choice struct {
	Text string `json:"text"`
}

type StoryTurn struct {
	NarrativeText    string   `json:"narrative_text"`
	NarrativeSummary string   `json:"narrative_summary"`
	Choices          []Choice `json:"choices"`
}
