package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_pypi_token(t *testing.T) {
	testContent := `
pypi-ASioasdfmkasdhusd89sd09ikasdasd
`
	pattern := PyPiAuthenticationToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "pypi-ASioasdfmkasdhusd89sd09ikasdasd", matches[0])
}
