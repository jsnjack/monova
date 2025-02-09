package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/tabwriter"
	"time"
)

// HistoryFilename is the name of the history file, which contains the mapping
// between the commit hash and the version number.
const HistoryFilename = ".monova.history"

const commitReadBatchSize = 1

var ErrNoCommitsLeft = fmt.Errorf("no commits left")

type History struct {
	path string // path to the repository, history file is stored in this directory
}

func (h *History) GetVersion(printHistory bool) (*PackageVersion, error) {
	started_at := time.Now()
	defer func() {
		DebugLogger.Printf("Calculating version took: %s\n", time.Since(started_at))
	}()
	// Load the latest version from the history file
	currentVersion, err := h.LoadVersionLog()
	if err != nil {
		DebugLogger.Printf("Failed to load version log: %v\n", err)
		currentVersion = &PackageVersion{}
	}

	i := 0
	commits := make([]*Commit, 0)

	// To track rebases, we need to check if the current commit is found in the git log.
	// If it is found, we reset currentVersion as everything will be calculated from scratch.
	currentCommitFound := false

	// Read all new commits from git
	for {
		commit, err := h.readCommitFromGit(i)
		if err != nil {
			if err == ErrNoCommitsLeft {
				DebugLogger.Printf("No more commits found. Exiting reading loop\n")
				break
			}
			return nil, fmt.Errorf("failed to read commit from git: %w", err)
		}

		if commit.CommitID == currentVersion.CommitID {
			DebugLogger.Printf("Current commit found: %s. Exiting reading loop\n", commit.CommitID)
			currentCommitFound = true
			break
		}

		commits = append(commits, commit)
		if commit.IsCheckpointCommit() && !printHistory {
			// If we find a checkpoint commit, we stop reading commits
			DebugLogger.Printf("Checkpoint commit found: %s. Exiting reading loop\n", commit.CommitID)
			break
		}
		i = i + commitReadBatchSize
	}

	if len(commits) == 0 {
		// No new commits, return the current version
		DebugLogger.Printf("No new commits found\n")
		return currentVersion, nil
	}

	// Reverse the commits, so that the first commit is the oldest one
	reverseCommits(commits)

	if !currentCommitFound {
		currentVersion = &PackageVersion{}
	}

	finalVersion, err := h.applyCommits(currentVersion, printHistory, commits...)
	if err != nil {
		return nil, err
	}

	err = h.SaveVersionLog(finalVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to save version log: %w", err)
	}
	return finalVersion, nil
}

func (h *History) applyCommits(versionIn *PackageVersion, printHistory bool, commits ...*Commit) (*PackageVersion, error) {
	DebugLogger.Printf("Applying %d commits to version %s\n", len(commits), versionIn.String())
	version := versionIn
	tabw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer func() {
		if printHistory {
			tabw.Flush()
		}
	}()
	for _, commit := range commits {
		// Apply the commit to the version
		changed := version.IncrementVersionFromMessage(commit)
		if printHistory {
			versionToPrint := ""
			if changed {
				versionToPrint = version.String()
			}
			fmt.Fprintf(tabw, "%s\t%s\t%s\n", versionToPrint, commit.CommitID, commit.Message)
		} else {
			DebugLogger.Printf("Commit %s: %s -> %s\n", commit.CommitID, commit.Message, version.String())
		}
	}
	return version, nil
}

func (h *History) readCommitFromGit(cursor int) (*Commit, error) {
	cmd := exec.Command("git", "log", "--format=%H %s", fmt.Sprintf("--max-count=%d", commitReadBatchSize), fmt.Sprintf("--skip=%d", cursor))
	cmd.Dir = filepath.Dir(h.path)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to read commit from git: %w", err)
	}

	// output is in the format: '<commitID> <message>'
	if len(output) == 0 {
		return nil, ErrNoCommitsLeft
	}

	commitStr := string(output)
	commit, err := NewCommit(commitStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create commit from string: %w", err)
	}
	return commit, nil
}

func (h *History) Reset() error {
	return os.Remove(h.path)
}

func (h *History) SaveVersionLog(version *PackageVersion) error {
	DebugLogger.Printf("Saving version log: %s\n", version.String())
	f, err := os.Create(h.path)
	if err != nil {
		return fmt.Errorf("failed to create history file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(version.HistoryString())
	if err != nil {
		return fmt.Errorf("failed to write version to history file: %w", err)
	}
	return nil
}

func (h *History) LoadVersionLog() (*PackageVersion, error) {
	DebugLogger.Printf("Loading version log\n")
	data, err := os.ReadFile(h.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read history file: %w", err)
	}

	version, err := NewPackageVersionFromVersionLog(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create version from history file: %w", err)
	}
	return version, nil
}

func NewHistory(path string) (*History, error) {
	DebugLogger.Printf("Creating history object for path: %s\n", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s does not exist", path)
	}

	// Verify that path is a directory and git repository
	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s is not a git repository", path)
	}

	history := &History{
		path: filepath.Join(path, HistoryFilename),
	}
	return history, nil
}

func reverseCommits(commits []*Commit) {
	for i, j := 0, len(commits)-1; i < j; i, j = i+1, j-1 {
		commits[i], commits[j] = commits[j], commits[i]
	}
}
