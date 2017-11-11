package main

import "testing"

func Test_getDataFromHistory_normal(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    1.10.2", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "f276e8b4d9d745e0914fde2f2eaba85e6c0de251" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 1 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 10 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 2 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}

func Test_getDataFromHistory_empty(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 0 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 0 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 0 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}

func Test_getDataFromHistory_bad_version(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    1.10", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 0 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 0 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 0 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}

func Test_getDataFromHistory_bad_version2(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "f276e8b4d9d745e0914fde2f2eaba85e6c0de251", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 0 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 0 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 0 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}

func Test_getDataFromHistory_bad_hash(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "f27 Add tests    1.10.2", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 0 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 0 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 0 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}

func Test_getDataFromHistory_bad_hash2(t *testing.T) {
	old := ReadLastLine
	defer func() {
		ReadLastLine = old
	}()
	ReadLastLine = func(path string) (string, error) {
		return "f276e8b4d9d74 5e0914fde2f2eaba85e6c0de251 Add tests    1.10.2", nil
	}
	version, commitID := getDataFromHistory("")
	if commitID != "" {
		t.Errorf("Uexpected result %s", commitID)
	}
	if version.Major != 0 {
		t.Errorf("Uexpected result %d", version.Major)
	}
	if version.Minor != 0 {
		t.Errorf("Uexpected result %d", version.Minor)
	}
	if version.Patch != 0 {
		t.Errorf("Uexpected result %d", version.Patch)
	}
}
