package main

import (
	"flag"
	"fmt"
	"os"
)

var version string

func main() {
	var oldVersion string

	infoFlag := flag.Bool("info", false, "Print old and new version")
	versionFlag := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

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
