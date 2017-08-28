package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var version string
var debugFlag *bool

func main() {
	var oldVersion string

	infoFlag := flag.Bool("info", false, "Print old and new version")
	versionFlag := flag.Bool("version", false, "Print version information")
	checkpointFlag := flag.Bool("checkpoint", false, "Create checkpoint")
	resetFlag := flag.Bool("reset", false, "Recalculate version")
	debugFlag = flag.Bool("debug", false, "Enable extra logging")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	cwd, _ := os.Getwd()

	// Remove .monova.cache file to recalculate the version
	if *resetFlag {
		cachePath := path.Join(cwd, cacheFilename)
		err := os.Remove(cachePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	repo, err := CreateRepository(cwd)
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
