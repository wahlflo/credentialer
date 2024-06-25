package public_keys

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_public_key_xml(t *testing.T) {
	testContent := `
<RSAKeyValue>
  <Modulus>qPfgaTEWEP3S9w0tgsicURfo+nLW09/0KfOPinhYZ4ouzU+3xC4pSlEp8Ut9FgL0AgqNslNaK34Kq+NZjO9DAQ==</Modulus>
  <Exponent>AQAB</Exponent>
</RSAKeyValue>
`
	pattern := PublicKeyInXml()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
}
