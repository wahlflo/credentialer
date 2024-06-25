package public_keys

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_public_key_pkcs1(t *testing.T) {
	testContent := `
-----BEGIN RSA PUBLIC KEY-----
MEgCQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnx
S30WAvQCCo2yU1orfgqr41mM70MBAgMBAAE=
-----END RSA PUBLIC KEY-----
`
	pattern := PkcsPublicKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
}

func Test_public_key_pkcs8(t *testing.T) {
	testContent := `
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf
9Cnzj4p4WGeKLs1Pt8QuKUpRKfFLfRYC9AIKjbJTWit+CqvjWYzvQwECAwEAAQ==
-----END PUBLIC KEY-----
`
	pattern := PkcsPublicKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
}
