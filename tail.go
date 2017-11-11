package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Reader interface
type Reader interface {
	ReadString(delim byte) (string, error)
}

// ReadLastLine reads last line with the content in the file
var ReadLastLine = func(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}
	breader := bufio.NewReader(file)
	return readLastLine(breader)
}

func readLastLine(breader Reader) (string, error) {
	lines := []string{}
	for {
		l, err := breader.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, l)
			break
		} else if err == nil {
			lines = append(lines, l)
		} else {
			return "", err
		}
	}
	return getLastMeaningfulLine(&lines)
}

func getLastMeaningfulLine(lines *[]string) (string, error) {
	for i := len(*lines) - 1; i >= 0; i-- {
		cleaned := strings.TrimSpace((*lines)[i])
		if len(cleaned) > 0 {
			return cleaned, nil
		}
	}
	return "", fmt.Errorf("Empty file")
}
