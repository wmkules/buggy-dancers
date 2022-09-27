package main

// promptStruct represents data about a record promptStruct.
type promptStruct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int64  `json:"votes"`
}

type ballotStruct struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Prompts     []promptStruct `json:"prompts"`
}

type voteStruct struct {
	BallotID string `json:"ballotID"`
	PromptID string `json:"promptID"`
}

var ballots = []ballotStruct{
	{ID: "1", Description: "Which is cuter?", Prompts: []promptStruct{
		{ID: "1", Name: "Puppies", Description: "", Votes: 3},
		{ID: "2", Name: "Kittens", Description: "", Votes: 4},
	}},
	{ID: "2", Description: "Dance style", Prompts: []promptStruct{
		{ID: "1", Name: "Swing", Description: "", Votes: 15},
		{ID: "2", Name: "Tappangucchi", Description: "", Votes: 5},
		{ID: "3", Name: "Break dance", Description: "", Votes: 10},
	}},
}
