package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToInt(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{
			input:    []byte{49, 50, 51, 52},
			expected: 1234,
		},
		{
			input:    []byte{49, 48, 48, 48, 48, 48, 48},
			expected: 1000000,
		},
		{
			input:    []byte{45, 49},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, BytesToInteger[int](tc.input))
	}
}
