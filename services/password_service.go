package services

import (
	"unicode"
)

// CalculateSteps calculates the number of steps required to make the password strong
func CalculateSteps(password string) int {
	n := len(password)

	// Check if the password contains at least one lowercase, one uppercase, and one digit
	hasLower := false
	hasUpper := false
	hasDigit := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	missingTypes := 3
	if hasLower {
		missingTypes--
	}
	if hasUpper {
		missingTypes--
	}
	if hasDigit {
		missingTypes--
	}

	// Calculate the number of repeating characters
	repeatSteps := 0
	i := 0
	for i < n {
		j := i
		for j < n && password[j] == password[i] {
			j++
		}
		repeatLen := j - i
		if repeatLen >= 3 {
			repeatSteps += repeatLen / 3
		}
		i = j
	}

	if n < 6 {
		// Not Enough password: add characters and address missing types
		return max(6-n, missingTypes)
	} else if n <= 20 {
		// Valid length: handle missing types and repetitions
		return max(repeatSteps, missingTypes)
	} else {
		// Too long: remove characters and then handle the rest
		deleteCount := n - 20
		// Reduce the repeatSteps by the number of deletions
		repeatSteps -= min(deleteCount, repeatSteps)
		// Total steps: deletions + maximum of remaining repetitions or missing types
		return deleteCount + max(repeatSteps, missingTypes)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
