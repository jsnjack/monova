package main

// Commit represebts a git commit
type Commit struct {
	ID      string
	Subject string
}

// CreateCommit creates a Commit object from the git log output
func CreateCommit(logStr string) (*Commit, error) {
	// Quotes is the first character
	commit := Commit{
		ID:      logStr[1:41],
		Subject: logStr[41:],
	}
	return &commit, nil
}
