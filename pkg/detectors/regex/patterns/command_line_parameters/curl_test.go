package command_line_parameters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HardcodedPasswordParameterInCurl_1(t *testing.T) {
	testContent := `
curl -u username:password https://example.com
`
	pattern := HardcodedPasswordParameterInCurl()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "curl -u username:password https://example.com", matches[0])
}
