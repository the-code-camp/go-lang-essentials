package main

import (
	"strings"
	"testing"
)

func Test_first_name_cannot_be_blank(t *testing.T) {
	err := addContact("", "", "")
	expected := "first name cannot be blank"
	if err.Error() != expected {
		t.Errorf("expected: %s, got %s", expected, err.Error())
	}
}

func Test_last_name_cannot_be_blank(t *testing.T) {
	err := addContact("firstname", "", "")
	expected := "last name cannot be blank"
	if err.Error() != expected {
		t.Errorf("expected: expected %s, got %s", expected, err.Error())
	}
}

func Test_phone_number_cannot_be_blank(t *testing.T) {
	err := addContact("firstname", "lastname", "")
	expected := "phone cannot be blank"
	if err.Error() != expected {
		t.Errorf("expected: expected %s, got %s", expected, err.Error())
	}
}

func Test_phone_number_length_should_be_more_than_10(t *testing.T) {
	err := addContact("firstname", "lastname", "phone")
	expected := "phone number should be of atleast 10 digits"
	if err.Error() != expected {
		t.Errorf("expected: expected %s, got %s", expected, err.Error())
	}
}

func Test_phone_number_should_start_with_zero(t *testing.T) {
	err := addContact("firstname", "lastname", "phone12345")
	expected := "phone number should start with zero"
	if err.Error() != expected {
		t.Errorf("expected: expected %s, got %s", expected, err.Error())
	}
}
func Test_phone_number_should_only_be_numeric(t *testing.T) {
	err := addContact("firstname", "lastname", "phone12345")
	expected := "phone number should only be numeric"

	if strings.Contains(err.Error(), expected) {
		t.Errorf("expected: expected %s, got %s", expected, err.Error())
	}
}
func Test_all_field_should_be_valid(t *testing.T) {
	err := addContact("firstname", "lastname", "0987654321")

	if err != nil {
		t.Errorf("expected: expected valid, got %s", err.Error())
	}
}
