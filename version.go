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
func (v *Version) Update(subject string) {
	if strings.Contains(subject, ":major:") {
		v.Major = v.Major + 1
	} else if strings.Contains(subject, ":minor:") {
		v.Minor = v.Minor + 1
	} else if strings.Contains(subject, ":patch:") {
		v.Patch = v.Patch + 1
	}
}
