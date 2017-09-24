package main

import "testing"

func TestProgressBar_Next_Count(t *testing.T) {
	pbar := ProgressBar{}
	pbar.Next()
	if pbar.count != 1 {
		t.Fatal("pbar.count was not increased")
	}
}

func TestProgressBar_Next_Return(t *testing.T) {
	pbar := ProgressBar{}
	result := pbar.Next()
	expected := "Analyzing commits: @1"
	if result != expected {
		t.Fatalf("Got %s, expected %s\n", result, expected)
	}
}
