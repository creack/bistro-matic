package parser

import (
	"fmt"
	"slices"
)

// ParseNumberBase parses a number in a given base and returns its integer value.
// Assume the input and base are valid.
// Supports utf8.
func ParseNumberBase(input, base string) (int, error) {
	if input == "" {
		return 0, fmt.Errorf("input number is empty")
	}
	if len(base) < 2 {
		return 0, fmt.Errorf("base must have at least two characters")
	}

	// Decode the base string into runes.
	// Needed for utf8 support.
	baseRunes := []rune(base)
	// Get the baseLen from the rune slice and not the string
	// as it would be wrong for utf8.
	baseLen := len(baseRunes)

	out := 0
	for _, c := range input {
		out *= baseLen
		idx := slices.Index(baseRunes, c)
		if idx == -1 {
			return 0, fmt.Errorf("character %c not found in base", c)
		}
		out += idx
	}
	return out, nil
}

// PutNumberBase converts a number to a string in a given base.
// Assume that the base is valid.
// Supports utf8.
func PutNumberBase(num int, base string) (string, error) {
	out := ""
	isNeg := num < 0
	if isNeg {
		num = -num
	}
	if len(base) < 2 {
		return "", fmt.Errorf("base must have at least two characters")
	}

	// Decode the base string into runes.
	// Needed for utf8 support.
	baseRunes := []rune(base)

	// Get the baseLen from the rune slice and not the string
	// as it would be wrong for utf8.
	baseLen := len(baseRunes)

	for num > 0 {
		out = string(baseRunes[num%baseLen]) + out
		num /= baseLen
	}
	if isNeg {
		out = "-" + out
	}
	return out, nil
}
