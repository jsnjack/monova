package cmd

import (
	"fmt"
	"strings"
)

// The length of the full commit ID hash in git is 40 characters
const commitIDLength = 40

type Commit struct {
	CommitID string
	Message  string
}

// NewCommit creates a new Commit object from a string
func NewCommit(str string) (*Commit, error) {
	// str is expected to be in the format of full commit ID followed by the commit message
	if len(str) < commitIDLength {
		return nil, fmt.Errorf("input string is too short")
	}

	commitID := str[:commitIDLength]
	message := str[commitIDLength:]

	if len(commitID) != commitIDLength {
		return nil, fmt.Errorf("invalid commit ID")
	}

	message = strings.TrimSpace(message)

	return &Commit{
		CommitID: commitID,
		Message:  message,
	}, nil
}
