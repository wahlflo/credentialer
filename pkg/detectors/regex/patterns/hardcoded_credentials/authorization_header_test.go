package hardcoded_credentials

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_authorization_http_header(t *testing.T) {
	testContent := `
Authorization=basic  Asdhuoiads:asduhaiusd
Authorization=bearer jksahnkjasd-asdadsads=
`
	pattern := AuthorizationHttpHeader()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "Asdhuoiads:asduhaiusd", matches[0])
	require.Equal(t, "jksahnkjasd-asdadsads=", matches[1])
}
