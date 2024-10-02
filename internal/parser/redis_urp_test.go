package parser

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRedisURP(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "*1\r\n$7\r\nCOMMAND\r\n",
			expected: "COMMAND",
		},
		{
			input:    "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n",
			expected: "GET foo",
		},
		{
			input:    "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
			expected: "SET foo bar",
		},
		{
			input:    "1\r\n$6\r\nFAILED\r\n",
			expected: "",
		},
	}

	for _, tc := range testCases {
		actual := ParseRedisURP(bufio.NewReader(bytes.NewReader([]byte(tc.input))))

		assert.Equal(t, tc.expected, actual)
	}
}
