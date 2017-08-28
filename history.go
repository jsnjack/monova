package main

import "os"
import "path"
import "fmt"

const historyFilename = ".monova.history"

// SaveHistory saves version history
func SaveHistory(line string) error {
	cwd, _ := os.Getwd()
	historyPath := path.Join(cwd, historyFilename)
	file, err := os.OpenFile(historyPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	if err != nil {
		return err
	}
	return nil
}
