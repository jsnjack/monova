package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var version string
var debugFlag *bool
var infoFlag *bool
var versionFlag *bool
var checkpointFlag *bool
var resetFlag *bool
var historyFlag *bool

func init() {
	infoFlag = flag.Bool("info", false, "Print old and new version")
	versionFlag = flag.Bool("version", false, "Print version information")
	checkpointFlag = flag.Bool("checkpoint", false, "Create checkpoint [version]")
	resetFlag = flag.Bool("reset", false, "Recalculate version")
	debugFlag = flag.Bool("debug", false, "Enable extra logging")
	historyFlag = flag.Bool("history", false, "Print version history")
	flag.Parse()
}

func main() {
	var oldVersion string

	// Print application version and exit
	if *versionFlag {
		if version == "" {
			fmt.Println("latest")
		} else {
			fmt.Println(version)
		}
		return
	}

	cwd, _ := os.Getwd()

	// Remove .monova.history file to recalculate the version
	if *resetFlag {
		historyPath := path.Join(cwd, historyFilename)
		err := os.Remove(historyPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Create repository instance
	repo, err := CreateRepository(cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save old version
	if *infoFlag {
		oldVersion = repo.GetVersion()
	}

	// Calculate new version
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

	if *historyFlag {
		err = PrintHistory()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
