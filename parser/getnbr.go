package parser

import (
	"fmt"
	"strings"
)

// ParseNumberBase parses a number in a given base and returns its integer value.
// Assume the input and base are valid.
// TODO: Support utf8.
func ParseNumberBase(input, base string) (int, error) {
	if input == "" {
		return 0, fmt.Errorf("input number is empty")
	}
	if base == "" {
		return 0, fmt.Errorf("base is empty")
	}
	baseLen := len(base)

	out := 0
	for _, c := range input {
		out *= baseLen
		idx := strings.IndexRune(base, c)
		if idx == -1 {
			return 0, fmt.Errorf("character %c not found in base", c)
		}
		out += idx
	}
	return out, nil
}
