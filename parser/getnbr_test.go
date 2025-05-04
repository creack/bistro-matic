package parser_test

import (
	"testing"

	"go.creack.net/bistro-matic/parser"
)

func TestConvertToInt(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		base     string
		expected int
		wantErr  bool
	}{
		{
			name:     "Base 10 - Regular number",
			number:   "123",
			base:     "0123456789",
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Base 2 - Binary",
			number:   "1010",
			base:     "01",
			expected: 10,
			wantErr:  false,
		},
		{
			name:     "Base 16 - Hexadecimal",
			number:   "1A",
			base:     "0123456789ABCDEF",
			expected: 26,
			wantErr:  false,
		},
		{
			name:     "Custom base - Letters",
			number:   "hello",
			base:     "abcdefghijklmnopqrstuvwxyz",
			expected: 3276872,
			wantErr:  false,
		},
		{
			name:     "Custom base - Symbols",
			number:   "@#$",
			base:     "!@#$%^&*()",
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Error - Empty base",
			number:   "123",
			base:     "",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Error - Empty number",
			number:   "",
			base:     "0123456789",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Error - Invalid character not in base",
			number:   "12A",
			base:     "0123456789",
			expected: 0,
			wantErr:  true,
		},
		// {
		// 	name:     "UTF-8 base",
		// 	number:   "你好世界",
		// 	base:     "你我他她它好世界的是不",
		// 	expected: 678,
		// 	wantErr:  false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseNumberBase(tt.number, tt.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNumberBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("ParseNumberBase() = %v, want %v", got, tt.expected)
			}
		})
	}
}
