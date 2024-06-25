package password_hashes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Sha256(t *testing.T) {
	testContent := `
$5$abc123$GI0iOPTIZol4fXe7fGtkWYVzY78oYOOAYiMglOkM4z7
`
	pattern := Sha256()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "$5$abc123$GI0iOPTIZol4fXe7fGtkWYVzY78oYOOAYiMglOkM4z7", matches[0])
}
