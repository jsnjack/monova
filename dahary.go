package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var oldVersion string

	infoFlag := flag.Bool("info", false, "Print old and new version")
	flag.Parse()

	path, _ := os.Getwd()
	repo, err := CreateRepository(path)
	if err != nil {
		panic(err)
	}

	if *infoFlag {
		oldVersion = repo.GetVersion()
	}

	newVersion, err := repo.UpdateVersion()
	if err != nil {
		panic(err)
	}

	if *infoFlag {
		fmt.Printf("%s -> %s\n", oldVersion, newVersion)
	} else {
		fmt.Println(newVersion)
	}
}
