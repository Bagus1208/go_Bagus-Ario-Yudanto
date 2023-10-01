package simple_unit_testing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	tests := []struct {
		name             string
		number1, number2 int
		expected         int
	}{
		{
			name:     "2 + 5",
			number1:  2,
			number2:  5,
			expected: 7,
		},
		{
			name:     "10 + 10",
			number1:  10,
			number2:  10,
			expected: 20,
		},
	}

	for _, test := range tests {
		result := Addition(test.number1, test.number2)
		require.Equal(t, test.expected, result, "Expected", test.expected)
	}
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		name             string
		number1, number2 int
		expected         int
	}{
		{
			name:     "20 - 5",
			number1:  20,
			number2:  5,
			expected: 15,
		},
		{
			name:     "10 + 10",
			number1:  10,
			number2:  10,
			expected: 0,
		},
	}

	for _, test := range tests {
		result := Subtraction(test.number1, test.number2)
		require.Equal(t, test.expected, result, "Expected", test.expected)
	}
}

func TestMultiple(t *testing.T) {
	tests := []struct {
		name             string
		number1, number2 int
		expected         int
	}{
		{
			name:     "20 x 5",
			number1:  20,
			number2:  5,
			expected: 100,
		},
		{
			name:     "15 + 5",
			number1:  15,
			number2:  5,
			expected: 75,
		},
	}

	for _, test := range tests {
		result := Multiple(test.number1, test.number2)
		require.Equal(t, test.expected, result, "Expected", test.expected)
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name             string
		number1, number2 float32
		expected         float32
	}{
		{
			name:     "20 / 5",
			number1:  20,
			number2:  5,
			expected: 4,
		},
		{
			name:     "14 / 2",
			number1:  14,
			number2:  2,
			expected: 7,
		},
	}

	for _, test := range tests {
		result := Division(test.number1, test.number2)
		require.Equal(t, test.expected, result, "Expected", test.expected)
	}
}
