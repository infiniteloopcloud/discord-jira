package jira

type Comment struct {
	Author User   `json:"author,omitempty" structs:"author,omitempty"`
	Body   string `json:"body,omitempty" structs:"body,omitempty"`
}
