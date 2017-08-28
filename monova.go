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
	checkpointFlag := flag.Bool("checkpoint", false, "Create checkpoint")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	path, _ := os.Getwd()
	repo, err := CreateRepository(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *infoFlag {
		oldVersion = repo.GetVersion()
	}

	newVersion, err := repo.UpdateVersion()
	if err != nil {
		fmt.Println(err)
		return
	}

	if *checkpointFlag {
		err = repo.CreateCheckpoint(flag.Args())
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *infoFlag {
		fmt.Printf("%s -> %s\n", oldVersion, newVersion)
		fmt.Printf("Commits inspected: %d\n", repo.commitCursor)
	} else {
		fmt.Println(newVersion)
	}
}
