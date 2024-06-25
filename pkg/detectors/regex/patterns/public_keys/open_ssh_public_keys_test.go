package public_keys

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_public_key_openssh(t *testing.T) {
	testContent := `
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41mM70MB phpseclib-generated-key
`
	pattern := OpenSSHPublicKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
}
