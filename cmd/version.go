package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type PackageVersion struct {
	CommitID string
	Major    int
	Minor    int
	Patch    int
}

func (pv *PackageVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", pv.Major, pv.Minor, pv.Patch)
}

func (pv *PackageVersion) HistoryString() string {
	return fmt.Sprintf("%s %s", pv.CommitID, pv.String())
}

func (pv *PackageVersion) ToCheckpointMessage() string {
	return fmt.Sprintf("%s %s %s", checkpointPrefix, pv.String(), checkpointSuffix)
}

func (pv *PackageVersion) IncrementVersionFromMessage(commit *Commit) bool {
	// Clear the message of any leading/trailing whitespace
	changed := false
	message := strings.TrimSpace(commit.Message)
	pv.CommitID = commit.CommitID

	// Check if the message is a checkpoint message
	if commit.IsCheckpointCommit() {
		versionStr := commit.GetVersionFromCheckpointMessage()
		var err error
		pv.Major, pv.Minor, pv.Patch, err = parseStringVersion(versionStr)
		if err != nil {
			DebugLogger.Printf("Failed to parse version from checkpoint message: %s\n", message)
			return changed
		}
		changed = true
		return changed
	}

	// Get the prefix of the message, which could be one of the version increment types
	msgPrefix := strings.Split(message, " ")[0]
	switch msgPrefix {
	case "p", ":p:", ":patch:":
		pv.Patch++
		changed = true
	case "m", ":m:", ":minor:":
		pv.Minor++
		pv.Patch = 0
		changed = true
	case "M", ":M:", ":major:":
		pv.Major++
		pv.Minor = 0
		pv.Patch = 0
		changed = true
	}
	return changed
}

func NewPackageVersionFromVersionLog(str string) (*PackageVersion, error) {
	// str is expected to be in the format of commitID major.minor.patch
	version := &PackageVersion{}
	commit, err := NewCommit(str)
	if err != nil {
		return nil, err
	}

	// Parse the version string
	version.Major, version.Minor, version.Patch, err = parseStringVersion(commit.Message)
	if err != nil {
		return nil, err
	}

	version.CommitID = commit.CommitID

	return version, nil
}

func parseStringVersion(str string) (int, int, int, error) {
	str = strings.TrimSpace(str)
	parts := strings.Split(str, ".")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid version format: %s", str)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid major version: %s", parts[0])
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid minor version: %s", parts[1])
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid patch version: %s", parts[2])
	}
	return major, minor, patch, nil
}
