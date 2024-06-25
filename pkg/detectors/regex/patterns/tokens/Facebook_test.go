package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_FacebookAccessToken_positive(t *testing.T) {
	testContent := `
access_token=EAACEdEose0cBANTYM3rlcZAFmfpoXBtde2qTs8RlXlu18JoLANzBZBAr1kRiBhS4ukg9sPAdZAIQgZCzGaOAA3fKrcZBvpdFihFZAqOdpCDGEGkyIoXOKxn3lZBFbUk6TKDKZCPemqjsu2HpIEUSsiRGZCSAeWzwsdAsC46nsxZALItwZDZD&lim
`
	pattern := FacebookAccessToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "EAACEdEose0cBANTYM3rlcZAFmfpoXBtde2qTs8RlXlu18JoLANzBZBAr1kRiBhS4ukg9sPAdZAIQgZCzGaOAA3fKrcZBvpdFihFZAqOdpCDGEGkyIoXOKxn3lZBFbUk6TKDKZCPemqjsu2HpIEUSsiRGZCSAeWzwsdAsC46nsxZALItwZDZD", matches[0])

}
