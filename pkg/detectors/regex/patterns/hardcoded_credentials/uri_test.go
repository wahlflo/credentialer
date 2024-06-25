package hardcoded_credentials

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_credentials_in_uri(t *testing.T) {
	testContent := `
https://test:abd@maps.googleapis.com/maps/api/js?key=YOUR_API_KEY&callback=initMap"
rdp://username:password@example.com
`
	pattern := CredentialsInUri()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "://test:abd@maps.googleapis.com", matches[0])
	require.Equal(t, "://username:password@example.com", matches[1])
}
