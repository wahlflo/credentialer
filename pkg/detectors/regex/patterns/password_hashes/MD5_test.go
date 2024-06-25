package password_hashes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MD5(t *testing.T) {
	testContent := `
$1$abc123$exkqoAsgtBy9b5sF0VJSW.
`
	pattern := MD5()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
}
