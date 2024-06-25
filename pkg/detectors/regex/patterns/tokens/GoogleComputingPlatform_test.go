package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_google_cloud_api_key(t *testing.T) {
	testContent := `
AIzaSyAxGxV2ApArc3rtM3WgYk0_wEaQx_wdOOg
`
	pattern := GoogleCloudApiKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "AIzaSyAxGxV2ApArc3rtM3WgYk0_wEaQx_wdOOg", matches[0])
}
