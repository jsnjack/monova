package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// BufferSize is the size of the buffer to read thelast line of the
// history file
var BufferSize = 512

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
	breader := createBuffer(file)
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
	if *debugFlag {
		fmt.Printf("Analysing %d lines from history file\n", len(*lines))
	}
	for i := len(*lines) - 1; i >= 0; i-- {
		cleaned := strings.TrimSpace((*lines)[i])
		if len(cleaned) > 0 {
			return cleaned, nil
		}
	}
	return "", fmt.Errorf("Empty file")
}

func createBuffer(file *os.File) Reader {
	breader := bufio.NewReaderSize(file, BufferSize)
	stat, _ := file.Stat()
	size := stat.Size()
	toDiscard := size - int64(BufferSize)
	if toDiscard > 0 {
		breader.Discard(int(toDiscard))
		if *debugFlag {
			fmt.Printf("Discarded %d bytes from history file\n", int(toDiscard))
		}
	}
	return breader
}
