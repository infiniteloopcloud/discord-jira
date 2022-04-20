package handler

import (
	"encoding/json"
	"strings"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/infiniteloopcloud/discord-jira/env"
	jira "github.com/infiniteloopcloud/discord-jira/jira"
)

type Issue struct {
	EventType string      `json:"webhookEvent"`
	Issue     *jira.Issue `json:"issue"`
	User      *jira.User  `json:"user"`
}

type Comment struct {
	EventType string        `json:"webhookEvent"`
	Issue     *jira.Issue   `json:"issue"`
	Comment   *jira.Comment `json:"comment"`
}

const (
	created = 0x90EE90
	deleted = 0xD10000
	updated = 0x0047AB
)

var channelName string

func Handle(eventType string, body []byte) (string, *discordgo.MessageEmbed, error) {
	switch eventType {
	case "jira:issue_created":
		return issueCreated(body)
	case "jira:issue_deleted":
		return issueDeleted(body)
	case "jira:issue_updated":
		return issueUpdated(body)
	case "comment_created":
		return commentCreated(body)
	case "comment_updated":
		return commentUpdated(body)
	case "comment_deleted":
		return commentDeleted(body)
	}
	return "", nil, nil
}

func issueCreated(body []byte) (string, *discordgo.MessageEmbed, error) {
	var issue Issue
	err := json.Unmarshal(body, &issue)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(issue.User.DisplayName, issue.User.AvatarUrls.Four8X48).
		SetTitle("["+issue.Issue.Fields.Project.Name+"] Issue created: "+issue.Issue.Fields.Summary).
		AddField("Issue type:", issue.Issue.Fields.Type.Name).
		AddField("Priority:", issue.Issue.Fields.Priority.Name).
		SetColor(created)

	if issue.Issue.Fields.Description != "" {
		if len(issue.Issue.Fields.Description) > 200 {
			message = message.AddField("Description", issue.Issue.Fields.Description[0:199]+"...")
		} else {
			message = message.AddField("Description", issue.Issue.Fields.Description)
		}
	}

	if issue.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + issue.Issue.Key)
	}

	channelName = strings.Replace(issue.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}

func issueDeleted(body []byte) (string, *discordgo.MessageEmbed, error) {
	var issue Issue
	err := json.Unmarshal(body, &issue)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(issue.User.DisplayName, issue.User.AvatarUrls.Four8X48).
		SetTitle("["+issue.Issue.Fields.Project.Name+"] Issue deleted: "+issue.Issue.Fields.Summary).
		AddField("Issue type:", issue.Issue.Fields.Type.Name).
		AddField("Priority:", issue.Issue.Fields.Priority.Name).
		SetColor(deleted)

	if issue.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + issue.Issue.Key)
	}

	channelName = strings.Replace(issue.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}

func issueUpdated(body []byte) (string, *discordgo.MessageEmbed, error) {
	var issue Issue
	err := json.Unmarshal(body, &issue)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(issue.User.DisplayName, issue.User.AvatarUrls.Four8X48).
		SetTitle("["+issue.Issue.Fields.Project.Name+"] Issue updated: "+issue.Issue.Fields.Summary).
		AddField("Issue type:", issue.Issue.Fields.Type.Name).
		AddField("Priority:", issue.Issue.Fields.Priority.Name).
		SetColor(updated)

	if issue.Issue.Fields.Description != "" {
		if len(issue.Issue.Fields.Description) > 200 {
			message = message.AddField("Description", issue.Issue.Fields.Description[0:199]+"...")
		} else {
			message = message.AddField("Description", issue.Issue.Fields.Description)
		}
	}

	if issue.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + issue.Issue.Key)
	}

	channelName = strings.Replace(issue.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}

func commentCreated(body []byte) (string, *discordgo.MessageEmbed, error) {
	var comment Comment
	err := json.Unmarshal(body, &comment)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(comment.Comment.Author.DisplayName, comment.Comment.Author.AvatarUrls.Four8X48).
		SetTitle("[" + comment.Issue.Fields.Project.Name + "] Comment created: " + comment.Issue.Fields.Summary).
		SetColor(created)

	if comment.Comment.Body != "" {
		if len(comment.Comment.Body) > 200 {
			message = message.AddField("Content", comment.Comment.Body[0:199]+"...")
		} else {
			message = message.AddField("Content", comment.Comment.Body)
		}
	}

	if comment.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + comment.Issue.Key)
	}

	channelName = strings.Replace(comment.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}

func commentDeleted(body []byte) (string, *discordgo.MessageEmbed, error) {
	var comment Comment
	err := json.Unmarshal(body, &comment)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(comment.Comment.Author.DisplayName, comment.Comment.Author.AvatarUrls.Four8X48).
		SetTitle("[" + comment.Issue.Fields.Project.Name + "] Comment deleted: " + comment.Issue.Fields.Summary).
		SetColor(deleted)

	if comment.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + comment.Issue.Key)
	}

	channelName = strings.Replace(comment.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}

func commentUpdated(body []byte) (string, *discordgo.MessageEmbed, error) {
	var comment Comment
	err := json.Unmarshal(body, &comment)
	if err != nil {
		return "", nil, err
	}

	message := embed.NewEmbed().
		SetAuthor(comment.Comment.Author.DisplayName, comment.Comment.Author.AvatarUrls.Four8X48).
		SetTitle("[" + comment.Issue.Fields.Project.Name + "] Comment updated: " + comment.Issue.Fields.Summary).
		SetColor(updated)

	if comment.Comment.Body != "" {
		if len(comment.Comment.Body) > 200 {
			message = message.AddField("Content", comment.Comment.Body[0:199]+"...")
		} else {
			message = message.AddField("Content", comment.Comment.Body)
		}
	}

	if comment.Issue.Key != "" {
		message = message.SetURL(env.Configuration().BaseURL + comment.Issue.Key)
	}

	channelName = strings.Replace(comment.Issue.Fields.Project.Name, " ", "-", 3)
	return channelName, message.MessageEmbed, nil
}
