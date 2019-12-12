package github

import "time"

// IssuesURL is the URL to the Github Issue Search API.
const IssuesURL = "https://api.github.com/search/issues"

// User represents a GitHub User.
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// Issue Represents a GitHub Issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // markdown format
}

// IssuesSearchResult represents a GitHub Issue Search Result.
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
