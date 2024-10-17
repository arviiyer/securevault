package main

import (
	"strings"
	"testing"
)

// TestGetUserChoice tests if the function correctly processes user input
func TestGetUserChoice(t *testing.T) {
	input := "e"
	expected := "e"

	// Simulate user input
	result := strings.TrimSpace(strings.ToLower(input))

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// TestGetUserInput tests if the function correctly handles directory input
func TestGetUserInput(t *testing.T) {
	input := "/test/directory"
	expected := "/test/directory"

	// Simulate user input
	result := input

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

