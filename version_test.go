package main

import (
	"fmt"
	"testing"
)

func TestVersion_Update_BadString(t *testing.T) {
	v := Version{}
	config := CreateConfig()
	err := v.Update("empty", config)
	if v.Major != 0 || v.Minor != 0 || v.Patch != 0 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 0.0.0\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_IncreaseMajor(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(":major: Commit", config)
	if v.Major != 2 || v.Minor != 0 || v.Patch != 0 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 2.0.0\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_IncreaseMinor(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(":minor: Commit", config)
	if v.Major != 1 || v.Minor != 2 || v.Patch != 0 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 1.2.0\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_IncreasePatch(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(":patch: Commit", config)
	if v.Major != 1 || v.Minor != 1 || v.Patch != 2 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 1.1.2\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_IncreasePatchP(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(":p: Commit", config)
	if v.Major != 1 || v.Minor != 1 || v.Patch != 2 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 1.1.2\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_IncreasePatchPMajorMultiple(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(":major: :minor: :patch: Commit", config)
	if v.Major != 2 || v.Minor != 0 || v.Patch != 0 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 2.0.0\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_Checkpoint(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(fmt.Sprintf("%s%s%s", checkpointPrefix, "3.2.1", checkpointSuffix), config)
	if v.Major != 3 || v.Minor != 2 || v.Patch != 1 {
		t.Fatalf(
			"Unexpected version value. Got %d.%d.%d, expected 3.2.1\n",
			v.Major, v.Minor, v.Patch,
		)
	}
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}

func TestVersion_Update_CheckpointBadVersion(t *testing.T) {
	v := Version{Major: 1, Minor: 1, Patch: 1}
	config := CreateConfig()
	err := v.Update(fmt.Sprintf("%s%s%s", checkpointPrefix, "boom", checkpointSuffix), config)
	if err == nil {
		t.Fatalf("Got no error\n")
	}
	if err.Error() != "Not valid checkpoint: boom" {
		t.Fatalf("Got unexpected error: %s\n", err.Error())
	}
}
