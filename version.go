package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents package version
type Version struct {
	Major int
	Minor int
	Patch int
}

// Update version based on the commit subject
func (v *Version) Update(subject string, config *Config) error {
	if strings.HasPrefix(subject, checkpointPrefix) && strings.HasSuffix(subject, checkpointSuffix) {
		checkpoint := subject[len(checkpointPrefix) : len(subject)-len(checkpointSuffix)]
		splits, err := SplitVersion(checkpoint)
		if err != nil {
			return err
		}
		v.Major = splits[0]
		v.Minor = splits[1]
		v.Patch = splits[2]
		return nil
	}

	for _, key := range config.MajorKeys {
		if strings.Contains(subject, key) {
			v.Major = v.Major + 1
			v.Minor = 0
			v.Patch = 0
			return nil
		}
	}

	for _, key := range config.MinorKeys {
		if strings.Contains(subject, key) {
			v.Minor = v.Minor + 1
			v.Patch = 0
			return nil
		}
	}

	for _, key := range config.PatchKeys {
		if strings.Contains(subject, key) {
			v.Patch = v.Patch + 1
			return nil
		}
	}
	return nil
}

// SplitVersion returns array of version [major, minor, patch]
func SplitVersion(checkpoint string) ([]int, error) {
	var numbers []int
	splits := strings.Split(checkpoint, ".")
	if len(splits) != 3 {
		return numbers, fmt.Errorf("Not valid checkpoint: %s", checkpoint)
	}
	for _, value := range splits {
		number, err := strconv.Atoi(value)
		if err != nil {
			return numbers, fmt.Errorf("Not valid checkpoint: %s", checkpoint)
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

// ExtractVersionFromString extracts version string, which is typically in the end
func ExtractVersionFromString(str *string) (string, error) {
	splitted := strings.SplitAfter(*str, " ")
	if len(splitted) > 0 {
		return splitted[len(splitted)-1], nil
	}
	return "", fmt.Errorf("Empty")
}
