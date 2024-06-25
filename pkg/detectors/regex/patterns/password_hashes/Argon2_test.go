package password_hashes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Argon2_positive(t *testing.T) {
	testContent := `
$argon2i$v=19$m=65536,t=2,p=1$c29tZXNhbHQ$AFcPXuKYWpT7cD1vTHHVlu2HiN91K5HCvM2mNeyTAMo
`
	pattern := Argon2()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "$argon2i$v=19$m=65536,t=2,p=1$c29tZXNhbHQ$AFcPXuKYWpT7cD1vTHHVlu2HiN91K5HCvM2mNeyTAMo", matches[0])
}
