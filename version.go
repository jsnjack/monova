package main

import (
	"strings"
)

// Version represents package version
type Version struct {
	Major int
	Minor int
	Patch int
}

// Update version based on the commit message
func (v *Version) Update(subject string, config *Config) {
	for _, key := range config.MajorKeys {
		if strings.Contains(subject, key) {
			v.Major = v.Major + 1
			v.Minor = 0
			v.Patch = 0
			return
		}
	}

	for _, key := range config.MinorKeys {
		if strings.Contains(subject, key) {
			v.Minor = v.Minor + 1
			v.Patch = 0
			return
		}
	}

	for _, key := range config.PatchKeys {
		if strings.Contains(subject, key) {
			v.Patch = v.Patch + 1
		}
		return
	}
	return
}
