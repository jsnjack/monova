package main

import "os"
import "path"
import "fmt"
import "strings"

const historyFilename = ".monova.history"

// SaveHistory saves version history
func SaveHistory(commit *Commit, version string) error {
	cwd, _ := os.Getwd()
	historyPath := path.Join(cwd, historyFilename)
	file, err := os.OpenFile(historyPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, generateHistoryLine(commit.ID, commit.Subject, version))
	if err != nil {
		return err
	}
	return nil
}

// Generates Histpry line. Trims subject if needed
func generateHistoryLine(commitID, subject, version string) string {
	var trimmed string
	switch {
	case len(subject) < 50:
		trimmed = subject + strings.Repeat(" ", 50-len(subject))
	default:
		trimmed = subject[:47] + "..."
	}
	line := fmt.Sprintf("%s %s %s", commitID, trimmed, version)
	return line
}
