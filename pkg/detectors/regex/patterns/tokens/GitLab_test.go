package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GitLabPersonalAccessToken_positive_1(t *testing.T) {
	testContent := `
glpat-sTcw8wx5hh7LpUusAkVh
`
	pattern := GitLabPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "glpat-sTcw8wx5hh7LpUusAkVh", matches[0])
}

func Test_GitLabPersonalAccessToken_positive_2(t *testing.T) {
	testContent := `
glpat-SNixgZ5e6NWeo1Wwga11
`
	pattern := GitLabPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "glpat-SNixgZ5e6NWeo1Wwga11", matches[0])
}

func Test_GitLabPersonalAccessToken_positive_3(t *testing.T) {
	testContent := `
"glpat-SNixgZ5e6NWeo1Wwga11"
"glpat-SNixgZ5e6NWeo1Wwga12"
`
	pattern := GitLabPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "glpat-SNixgZ5e6NWeo1Wwga11", matches[0])
	require.Equal(t, "glpat-SNixgZ5e6NWeo1Wwga12", matches[1])
}

func Test_GitLabPersonalAccessToken_negative(t *testing.T) {
	testContent := `
glpat-sTcw8wx5hh7LpUusAkVhTooLONG
glpat-SNixgZ5e6NWeo1Wwgas
`
	pattern := GitLabPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}
