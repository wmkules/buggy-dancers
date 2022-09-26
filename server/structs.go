package main

// prompt represents data about a record prompt.
type prompt struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int64  `json:"votes"`
}

type ballot struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Prompts     []prompt `json:"prompts"`
}

type vote struct {
	BallotID string `json:"ballotID"`
	PromptID string `json:"promptID"`
}

var ballots = []ballot{
	{ID: "1", Description: "Which is cuter?", Prompts: []prompt{
		{ID: "1", Name: "Puppies", Description: "", Votes: 3},
		{ID: "2", Name: "Kittens", Description: "", Votes: 4},
	}},
	{ID: "2", Description: "Dance style", Prompts: []prompt{
		{ID: "1", Name: "Swing", Description: "", Votes: 15},
		{ID: "2", Name: "Tappangucchi", Description: "", Votes: 5},
		{ID: "3", Name: "Break dance", Description: "", Votes: 10},
	}},
}
