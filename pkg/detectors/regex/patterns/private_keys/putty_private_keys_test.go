package private_keys

import (
	"github.com/stretchr/testify/require"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"testing"
)

func Test_private_key_putty_version2(t *testing.T) {
	testContent := `
PuTTY-User-Key-File-2: ssh-rsa
Encryption: none
Comment: phpseclib-generated-key
Public-Lines: 2
AAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp
84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41mM70MB
Private-Lines: 4
AAAAQCCS4sQctqqRaEuA0pqBJqN6hstLes9PQCCbR/uTnkVdW3vjeHA2CSn3xsw2
vPL0BDWYkZtBkaumvhzxkDHdpE0AAAAhAKMSvzIBnni7ot/OSie2TmJLY4SwTQAe
vXysE2RbFDYdAAAAIQEJQRpFCcydunv2bENcN/oBTRw39E8GNv2pIcNxZkcbNQAA
ACATrP+Toj4KE3Usu23BfSBqUhPGYBis4GEFWXjfe2BNNA==
Private-MAC: bc712a70870b4b8ddf120530f02b9068e782a21a
`
	pattern := PuttyPrivateKeys()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	puttyKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityHigh)
}

func Test_private_key_putty_version2_encrypted(t *testing.T) {
	testContent := `
PuTTY-User-Key-File-2: ssh-rsa
Encryption: aes256-cbc
Comment: phpseclib-generated-key
Public-Lines: 2
AAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp
84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41mM70MB
Private-Lines: 4
vdqfIladR4JIsN6wmmfJ9rt+PzuY+sZVP/vbsiOODeU24BYGj5arK/qjC2Bsr8vU
h/bkkK9AVqzd5sPaMzQ3HPya+ogEDoTKTr3SKg+twjItQb7q2gHwIvebPw67i8HN
hL+DmVZ2cJ1BXDHt79wJMQApmMlpRhJs0QziWhu1nTfnb8dPcC4B1RKCLPvtv+Iw
AjomU6mTfZs3ZVrxkH8e50q7cbkxQinQ9Su/9jYvtjIjMT6C0jgRUQUrQIHGFKGX
Private-MAC: 9b86c10fa3325ed53cee0d283a3f71791355acd1
`
	pattern := PuttyPrivateKeys()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	puttyKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityInformative)
}

func Test_private_key_putty_version3(t *testing.T) {
	testContent := `
PuTTY-User-Key-File-3: ssh-rsa
Encryption: none
Comment: phpseclib-generated-key
Public-Lines: 2
AAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp
84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41mM70MB
Private-Lines: 4
AAAAQCCS4sQctqqRaEuA0pqBJqN6hstLes9PQCCbR/uTnkVdW3vjeHA2CSn3xsw2
vPL0BDWYkZtBkaumvhzxkDHdpE0AAAAhAKMSvzIBnni7ot/OSie2TmJLY4SwTQAe
vXysE2RbFDYdAAAAIQEJQRpFCcydunv2bENcN/oBTRw39E8GNv2pIcNxZkcbNQAA
ACATrP+Toj4KE3Usu23BfSBqUhPGYBis4GEFWXjfe2BNNA==
Private-MAC: 53ba974a4a5f8ac69eb526fd0556fe1a5ccf654216d261af04aca910967b2204
`
	pattern := PuttyPrivateKeys()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	puttyKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityHigh)
}

func Test_private_key_putty_version3_encrypted(t *testing.T) {
	testContent := `
PuTTY-User-Key-File-3: ssh-rsa
Encryption: aes256-cbc
Comment: phpseclib-generated-key
Public-Lines: 2
AAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp
84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41mM70MB
Key-Derivation: Argon2id
Argon2-Memory: 8192
Argon2-Passes: 13
Argon2-Parallelism: 1
Argon2-Salt: 1ca5840ce3a82760ae2390f493adddb6
Private-Lines: 4
SD50LGS5iJOv4EY6HFbYoepp59L9+ACiMfMPfw8PxrWrsnNDOIyQQllCppzz6Z89
0ph7orDXyda47q7GzG6hFKfMtLC9p0o7ri8LHAnbj8uRHtMF2L0CWmimYbA8QSA+
GKUpm9yXMwmwEs6hsGJSCBV+T4jz573Jw2W1FnQRzFDwqaae9UWMEPfhQEwTTAbC
U1ljnk8H7g0JR8A58RBuqCxL95Azko2jXx/iyX1V94TCG3QfV/hc1fmX+XHHb7ly
Private-MAC: 99e41bb362c13774f506b0bfd5998771c4b61c14b8176dabbdbb631871b03652
`
	pattern := PuttyPrivateKeys()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	puttyKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityInformative)
}

func puttyKeysEnsureProperPriority(t *testing.T, testData string, priority interfaces.FindingPriority) {
	finding := interfaces.FindingInstance{
		Value:    testData,
		Priority: interfaces.FindingPriorityHigh,
	}

	require.Equal(t, priority, puttyPrivateKeyQualityCheck(finding, nil, nil).GetFindingPriority())
}
