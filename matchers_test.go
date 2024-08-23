package matchers

import (
	"context"
	"testing"
	"time"
)

func TestAnyContextMatcher(t *testing.T) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tests := []struct {
		name     string
		value    any
		expected bool
	}{
		{
			name:     "test context background",
			value:    context.Background(),
			expected: true,
		},
		{
			name:     "test context with value",
			value:    context.WithValue(context.Background(), "key", "value"),
			expected: true,
		},
		{
			name:     "test context with timeout",
			value:    ctxWithTimeout,
			expected: true,
		},
		{
			name:     "nil",
			value:    nil,
			expected: false,
		},
		{
			name:     "int",
			value:    120,
			expected: false,
		},
	}

	matcher := AnyContext()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isMatched := matcher.Matches(tt.value)
			if tt.expected != isMatched {
				t.Errorf(`"%v": got %t, want %t.`, tt.value, isMatched, tt.expected)
			}
		})
	}
}
