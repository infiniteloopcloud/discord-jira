package jira

type Issue struct {
	Key    string      `json:"key,omitempty" structs:"key,omitempty"`
	Fields IssueFields `json:"fields,omitempty" structs:"fields,omitempty"`
}

type IssueFields struct {
	Type     IssueType `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Project  Project   `json:"project,omitempty" structs:"project,omitempty"`
	Priority Priority  `json:"priority,omitempty" structs:"priority,omitempty"`
	Summary  string    `json:"summary,omitempty" structs:"summary,omitempty"`
}

type IssueType struct {
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

type Project struct {
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

type Priority struct {
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}
