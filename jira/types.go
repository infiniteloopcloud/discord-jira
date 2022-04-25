package jira

type IssueWrapper struct {
	EventType string `json:"webhookEvent"`
	Issue     *Issue `json:"issue"`
	User      *User  `json:"user"`
}

type CommentWrapper struct {
	EventType string   `json:"webhookEvent"`
	Issue     *Issue   `json:"issue"`
	Comment   *Comment `json:"comment"`
}
