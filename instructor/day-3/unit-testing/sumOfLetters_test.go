package main

import (
	"testing"
)

// 0, 1, multiple
func Test_sumOfLetters_should_return_0_when_empty(t *testing.T) {
	result := sumOfLetters("")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
func Test_sumOfLetters_should_return_1_when_input_is_a(t *testing.T) {
	result := sumOfLetters("a")
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
func Test_sumOfLetters_should_return_26_when_input_is_z(t *testing.T) {
	result := sumOfLetters("z")
	if result != 26 {
		t.Errorf("Expected 26, got %d", result)
	}
}

func Test_sumOfLetters_should_return_6_when_input_is_cab(t *testing.T) {
	result := sumOfLetters("cab")
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
func Test_sumOfLetters_should_return_100_when_input_is_excellent(t *testing.T) {
	result := sumOfLetters("excellent")
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}
}
