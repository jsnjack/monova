package main

// Cache is an object that stores information about the latest version
type Cache struct {
	Version  *Version
	CommitID string
}

// CreateCache creates Cache instance
func CreateCache(path string) (*Cache, error) {
	version, commitID := GetDataFromHistory(path)
	cache := Cache{Version: version, CommitID: commitID}
	return &cache, nil
}

// GetDataFromHistory returns Version object and commitid from the history line
func GetDataFromHistory(path string) (*Version, string) {
	var version Version
	versionLine, err := ReadLastLine(path)
	if err != nil {
		return &version, ""
	}
	versionStr, err := ExtractVersionFromString(&versionLine)
	if err != nil {
		return &version, ""
	}
	versionList, err := SplitVersion(versionStr)
	if err != nil {
		return &version, ""
	}

	commitID, err := GetCommitID(&versionLine)
	if err != nil {
		return &version, ""
	}

	version.Major = versionList[0]
	version.Minor = versionList[1]
	version.Patch = versionList[2]
	return &version, commitID
}
