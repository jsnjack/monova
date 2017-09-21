package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

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

// PrintHistory prints version history. Use it to find commits related to the
// specific version or vice versa
func PrintHistory() error {
	cwd, _ := os.Getwd()
	historyPath := path.Join(cwd, historyFilename)
	file, err := os.OpenFile(historyPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		fmt.Print(line)
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return err
	}
	return nil
}
