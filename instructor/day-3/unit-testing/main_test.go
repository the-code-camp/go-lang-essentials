package main

import "testing"

func Test_should_return_greeting(t *testing.T) {
	got := greeting()
	expected := "Hello"
	if got != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}

func greeting() string {
	return "Hello"
}
