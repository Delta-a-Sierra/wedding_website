package entities_test

import (
	"testing"
	"time"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateCountDownFromDateTime(t *testing.T) {
	tests := map[string]struct {
		input    time.Time
		expected entities.Countdown
	}{
		"FutureTime simple": {
			input:    time.Now().Add(24 * time.Hour),
			expected: entities.Countdown{Days: 0, Hours: 23, Minutes: 59},
		},
		"PastTime simple": {
			input:    time.Now().Add(-24 * time.Hour),
			expected: entities.Countdown{Days: -1, Hours: 0, Minutes: 0},
		},
		"FutureTime moderate": {
			input:    time.Now().Add(24 * time.Hour).Add(5 * time.Minute),
			expected: entities.Countdown{Days: 1, Hours: 0, Minutes: 4},
		},
		"PastTime moderate": {
			input:    time.Now().Add(-24 * time.Hour).Add(-5 * time.Minute),
			expected: entities.Countdown{Days: -1, Hours: 0, Minutes: -5},
		},
		"FutureTime complete": {
			input:    time.Now().Add(24 * time.Hour).Add(3 * time.Hour).Add(5 * time.Minute),
			expected: entities.Countdown{Days: 1, Hours: 3, Minutes: 4},
		},
		"PastTime complete": {
			input:    time.Now().Add(-24 * time.Hour).Add(-3 * time.Hour).Add(-5 * time.Minute),
			expected: entities.Countdown{Days: -1, Hours: -3, Minutes: -5},
		},
		"FutureTime complete 2": {
			input:    time.Now().Add(120 * time.Hour).Add(8 * time.Hour).Add(32 * time.Minute).Add(40 * time.Second),
			expected: entities.Countdown{Days: 5, Hours: 8, Minutes: 32},
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := entities.CreateCountDownFromDateTime(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCountDown_ZeroOutNegativeValues(t *testing.T) {
	tests := map[string]struct {
		input    entities.Countdown
		expected entities.Countdown
	}{
		"no negatives": {
			input:    entities.Countdown{Days: 3, Hours: 1, Minutes: 2},
			expected: entities.Countdown{Days: 3, Hours: 1, Minutes: 2},
		},
		"all negatives": {
			input:    entities.Countdown{Days: -3, Hours: -1, Minutes: -2},
			expected: entities.Countdown{Days: 0, Hours: 0, Minutes: 0},
		},
		"mixed": {
			input:    entities.Countdown{Days: 3, Hours: -1, Minutes: 2},
			expected: entities.Countdown{Days: 3, Hours: 0, Minutes: 2},
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			tc.input.ZeroOutMinusValues()
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
