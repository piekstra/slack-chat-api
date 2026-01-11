package messages

import (
	"testing"
)

func TestFormatTimestamp(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectSame bool // If true, expect output equals input (for invalid inputs)
	}{
		{
			name:  "standard timestamp",
			input: "1704067200.123456",
		},
		{
			name:  "timestamp without decimal",
			input: "1704067200",
		},
		{
			name:       "empty string",
			input:      "",
			expectSame: true,
		},
		{
			name:       "invalid timestamp",
			input:      "not-a-timestamp",
			expectSame: true,
		},
		{
			name:  "timestamp with extra precision",
			input: "1704067200.123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatTimestamp(tt.input)
			if tt.expectSame {
				if result != tt.input {
					t.Errorf("formatTimestamp(%q) = %q, expected %q", tt.input, result, tt.input)
				}
			} else {
				// For valid timestamps, check the format is correct (YYYY-MM-DD HH:MM)
				if len(result) != 16 {
					t.Errorf("formatTimestamp(%q) = %q, expected 16-char format YYYY-MM-DD HH:MM", tt.input, result)
				}
				// Check it contains expected delimiters
				if result[4] != '-' || result[7] != '-' || result[10] != ' ' || result[13] != ':' {
					t.Errorf("formatTimestamp(%q) = %q, format doesn't match YYYY-MM-DD HH:MM", tt.input, result)
				}
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{
			name:     "short string no truncation",
			input:    "Hello",
			maxLen:   10,
			expected: "Hello",
		},
		{
			name:     "exact length",
			input:    "Hello",
			maxLen:   5,
			expected: "Hello",
		},
		{
			name:     "truncation needed",
			input:    "Hello World!",
			maxLen:   8,
			expected: "Hello...",
		},
		{
			name:     "newlines converted to spaces",
			input:    "Hello\nWorld",
			maxLen:   20,
			expected: "Hello World",
		},
		{
			name:     "multiple newlines",
			input:    "Line1\nLine2\nLine3",
			maxLen:   20,
			expected: "Line1 Line2 Line3",
		},
		{
			name:     "truncation with newlines",
			input:    "Hello\nWorld\nFoo\nBar",
			maxLen:   10,
			expected: "Hello W...",
		},
		{
			name:     "empty string",
			input:    "",
			maxLen:   10,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncate(tt.input, tt.maxLen)
			if result != tt.expected {
				t.Errorf("truncate(%q, %d) = %q, expected %q", tt.input, tt.maxLen, result, tt.expected)
			}
		})
	}
}

func TestBuildDefaultBlocks(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "simple text",
			text: "Hello World",
		},
		{
			name: "markdown text",
			text: "*bold* _italic_ ~strike~",
		},
		{
			name: "empty text",
			text: "",
		},
		{
			name: "text with special characters",
			text: "Hello <@U123> in #general",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildDefaultBlocks(tt.text)

			if len(result) != 1 {
				t.Fatalf("expected 1 block, got %d", len(result))
			}

			block, ok := result[0].(map[string]interface{})
			if !ok {
				t.Fatal("expected block to be map[string]interface{}")
			}

			if block["type"] != "section" {
				t.Errorf("expected block type 'section', got %v", block["type"])
			}

			textObj, ok := block["text"].(map[string]interface{})
			if !ok {
				t.Fatal("expected text to be map[string]interface{}")
			}

			if textObj["type"] != "mrkdwn" {
				t.Errorf("expected text type 'mrkdwn', got %v", textObj["type"])
			}

			if textObj["text"] != tt.text {
				t.Errorf("expected text %q, got %v", tt.text, textObj["text"])
			}
		})
	}
}
