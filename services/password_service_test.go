package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		password   string
		numOfSteps int
	}{
		{"aA1", 3},                      // Needs 3 more characters to be valid
		{"1445D1cd", 0},                 // Already valid
		{"abcdef", 2},                   // Needs to replace missing uppercase and digit
		{"ABCDEF", 2},                   // Needs to replace missing lowercase and digit
		{"1234567", 2},                  // Needs to replace missing uppercase
		{"12345678", 2},                 // Needs to replace missing uppercase
		{"1111111111", 5},               // Needs 3 replacements and 2 additions for diversity
		{"aaaaaaaaaaaaaaa", 7},          // Needs 5 replacements and 2 additions for diversity
		{"aB1CdefGhIjKlmNopQrStUvW", 4}, // Too long, needs 4 deletions
		{"", 6},                         // Too short, needs 6 additions
		{"a", 5},                        // Too short, needs 5 additions
		{"aA1bB2cC3", 0},                // Just right with variety
		{"aaa111bbb", 4},                // Needs 3 replacements and 1 addition for diversity
		{"aaab", 5},                     // Needs 1 replacement, 2 additions, 2 more replacements for diversity
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			got := CalculateSteps(tt.password)
			assert.Equal(t, tt.numOfSteps, got, "CalculateSteps() = %v, want %v", got, tt.numOfSteps)
		})
	}
}
