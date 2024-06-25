package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GitHubPersonalAccessToken_positive(t *testing.T) {
	testContent := `
ghp_6OueNgi8nQJ2GSekRmNOqlxZgPtkrY2dkxGS
`
	pattern := GitHubPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ghp_6OueNgi8nQJ2GSekRmNOqlxZgPtkrY2dkxGS", matches[0])
}

func Test_GitHubPersonalAccessToken_negative(t *testing.T) {
	testContent := `
GHP_EXAMPLE_TEST_1
`
	pattern := GitHubPersonalAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}

func Test_GitHubOAuthAccessToken_positive(t *testing.T) {
	testContent := `
access_token=gho_16C7e42F292c6912E7710c838347Ae178B4a&scope=repo%2Cgist&token_type=bearer
`
	pattern := GitHubOAuthAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "gho_16C7e42F292c6912E7710c838347Ae178B4a", matches[0])
}
