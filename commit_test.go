package main

import "testing"

func Test_CreateCommit_good(t *testing.T) {
	str := "'f3b13c7b53737be455eafe3541c1561d4d3b599d:p: Rename getDataFromHistroy function'\n"
	commit, err := CreateCommit(str)
	if err != nil {
		t.Error(err)
		return
	}
	expected1 := "f3b13c7b53737be455eafe3541c1561d4d3b599d"
	if commit.ID != expected1 {
		t.Errorf("Expected %s, got %s \n", expected1, commit.ID)
	}

	expected2 := ":p: Rename getDataFromHistroy function"
	if commit.Subject != expected2 {
		t.Errorf("Expected %s, got %s \n", expected2, commit.Subject)
	}
}

func Test_CreateCommit_bad(t *testing.T) {
	str := "'f'\n"
	_, err := CreateCommit(str)
	if err == nil {
		t.Errorf("Expected error")
	}
}
