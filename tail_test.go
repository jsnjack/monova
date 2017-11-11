package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_readLastLine_one(t *testing.T) {
	data := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2"
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	result, err := readLastLine(breader)
	if result != data {
		t.Errorf("Expected %s, got %s", data, result)
	}
	if err != nil {
		t.Error(err)
	}
}

func Test_readLastLine_one_with_endline(t *testing.T) {
	data := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2\n"
	expected := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2"
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	result, err := readLastLine(breader)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	if err != nil {
		t.Error(err)
	}
}

func Test_readLastLine_one_with_2endline(t *testing.T) {
	data := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2\n\n"
	expected := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2"
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	result, err := readLastLine(breader)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	if err != nil {
		t.Error(err)
	}
}

func Test_readLastLine_one_with_2line(t *testing.T) {
	data := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.1\nf276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2"
	expected := "f276e8b4d9d745e0914fde2f2eaba85e6c0de251 Add tests    0.10.2"
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	result, err := readLastLine(breader)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	if err != nil {
		t.Error(err)
	}
}

func Test_readLastLine_newline_only(t *testing.T) {
	data := "\n"
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	_, err := readLastLine(breader)
	if err == nil {
		t.Errorf("Expected error\n")
		return
	}
	if err.Error() != "Empty file" {
		t.Error(err)
	}
}

func Test_readLastLine_space_only(t *testing.T) {
	data := " "
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	_, err := readLastLine(breader)
	if err == nil {
		t.Errorf("Expected error\n")
		return
	}
	if err.Error() != "Empty file" {
		t.Error(err)
	}
}

func Test_readLastLine_empty(t *testing.T) {
	data := ""
	reader := strings.NewReader(data)
	breader := bufio.NewReader(reader)
	_, err := readLastLine(breader)
	if err == nil {
		t.Errorf("Expected error\n")
		return
	}
	if err.Error() != "Empty file" {
		t.Error(err)
	}
}
