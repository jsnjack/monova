package main

import "strings"
import "fmt"

// Commit represebts a git commit
type Commit struct {
	ID      string
	Subject string
}

// CreateCommit creates a Commit object from the git log output
func CreateCommit(logStr string) (*Commit, error) {
	// Quotes is the first character
	cleaned := strings.Trim(logStr, " \n'")
	if len(cleaned) < 41 {
		return nil, fmt.Errorf("Not valid log string from git: %s", logStr)
	}
	commit := Commit{
		ID:      cleaned[:40],
		Subject: cleaned[40:],
	}
	return &commit, nil
}
