package models

type Problem struct {
	ContestID int    `bson:"contest_id",json:"contest_id"`
	Index     string `bson:"index",json:"index"`
}

type Submission struct {
	ID      int     `bson:"id",json:"id"`
	Handle  string  `bson:"handle,omitempty",bson:"handle,omitempty"`
	Lang    string  `bson:"lang,omitempty",json:"lang,omitempty"`
	Verdict string  `bson:"verdict,omitempty",json:"verdict,omitempty"`
	Problem Problem `bson:"problem",json:"problem"`
}

type Testcase struct {
	Input             *string `bson:"input,omitempty",json:"input,omitempty"`
	JuryOutput        *string `bson:"jury_output,omitempty",json:"jury_output,omitempty"`
	ParticipantOutput *string `bson:"participant_output,omitempty",json:"participant_output,omitempty"`
}

type Ticket struct {
	TicketID   int         `bson:"ticket_id",json:"ticket_id"`
	Type       string      `bson:"type,omitempty",json:"type,omitempty"`
	Progress   string      `bson:"progress,omitempty",json:"progress,omitempty"`
	Verdict    string      `bson:"verdict,omitempty",json:"verdict,omitempty"`
	Problem    Problem     `bson:"problem",json:"problem"`
	Submission *Submission `bson:"submission,omitempty",json:"submission,omitempty"`
	Testcase   *Testcase   `bson:"testcase,omitempty",json:"testcase,omitempty"`
	Parameters string      `bson:"parameters,omitempty",json:"parameters,omitempty"`
}
