package main

import (
	"testing"
)

func Test_should_return_the_sum_of_two_numbers(t *testing.T) {
	// Arrange
	expected := 30

	// Act
	actual := sum(10, 20)

	// Assert
	if actual != expected {
		t.Errorf("expected 30, recevied: %d", actual)
	}
}

func Test_sum_functionality_table_driven_test(t *testing.T) {
	tests := []struct {
		testName string
		x        int
		y        int
		expected int
	}{
		{
			"testing with 10, 20",
			10,
			20,
			30,
		},
		{
			"testing with 20, 20",
			20,
			20,
			40,
		},
		{
			"testing with 10, 60",
			10,
			60,
			70,
		},
	}

	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {

			if actual := sum(tt.x, tt.y); actual != tt.expected {
				t.Errorf("Sum failed, expected = %d, actual=%d", tt.expected, actual)
			}

		})
	}
}

func sum(x, y int) int {
	return x + y
}
