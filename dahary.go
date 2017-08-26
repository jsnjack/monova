package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Getwd()
	repo, err := CreateRepository(path)
	if err != nil {
		panic(err)
	}
	newVersion, err := repo.UpdateVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVersion)
}
