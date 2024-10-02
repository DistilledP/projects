package parser

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/DistilledP/lungfish/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestParseRedisURP(t *testing.T) {
	testCases := []struct {
		input    string
		expected types.Command
	}{
		{
			input:    "*1\r\n$7\r\nCOMMAND\r\n",
			expected: types.Command{Name: "COMMAND", Args: []string{}},
		},
		{
			input:    "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n",
			expected: types.Command{Name: "GET", Args: []string{"foo"}},
		},
		{
			input:    "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
			expected: types.Command{Name: "SET", Args: []string{"foo", "bar"}},
		},
		{
			input:    "1\r\n$6\r\nFAILED\r\n",
			expected: types.Command{},
		},
	}

	for _, tc := range testCases {
		actual := ParseRedisURP(bufio.NewReader(bytes.NewReader([]byte(tc.input))))

		assert.Equal(t, tc.expected, actual)
	}
}
