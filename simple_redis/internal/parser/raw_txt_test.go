package parser

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DistilledP/projects/simple_redis/internal/types"
)

func TestRaw(t *testing.T) {
	testCases := []struct {
		input    string
		expected types.Command
	}{
		{
			input:    "COMMAND\r\n",
			expected: types.Command{Name: "command", Args: []string{}},
		},
		{
			input:    "GET foo\r\n",
			expected: types.Command{Name: "get", Args: []string{"foo"}},
		},
		{
			input:    "SET foo bar\r\n",
			expected: types.Command{Name: "set", Args: []string{"foo", "bar"}},
		},
		{
			input:    "\r\n",
			expected: types.Command{},
		},
	}

	for _, tc := range testCases {
		actual := ParseRaw(bufio.NewReader(bytes.NewReader([]byte(tc.input))))

		assert.Equal(t, tc.expected, actual)
	}
}
