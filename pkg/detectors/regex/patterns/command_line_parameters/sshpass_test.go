package command_line_parameters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sshpass_positive_1(t *testing.T) {
	testContent := `
sshpass -p 'MY PASSWORT' ssh dein_benutzername@dein_server
`
	pattern := HardcodedPasswordParameterInSshpass()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "sshpass -p 'MY PASSWORT'", matches[0])
}

func Test_sshpass_positive_2(t *testing.T) {
	testContent := `
sshpass -p "MY PASSWORT" ssh dein_benutzername@dein_server
`
	pattern := HardcodedPasswordParameterInSshpass()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "sshpass -p \"MY PASSWORT\"", matches[0])
}

func Test_sshpass_positive_3(t *testing.T) {
	testContent := `
sshpass -p MY_PASSWORT ssh dein_benutzername@dein_server
`
	pattern := HardcodedPasswordParameterInSshpass()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "sshpass -p MY_PASSWORT", matches[0])
}
