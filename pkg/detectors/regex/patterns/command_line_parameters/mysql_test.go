package command_line_parameters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HardcodedPasswordParameterInMySqlCli_1(t *testing.T) {
	testContent := `
mysql -u username -p'password' -h hostname database_name
`
	pattern := HardcodedPasswordParameterInMySqlCli()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "mysql -u username -p'password'", matches[0])
}

func Test_HardcodedPasswordParameterInMySqlCli_2(t *testing.T) {
	testContent := `
mysql -u username -p"password" -h hostname database_name
`
	pattern := HardcodedPasswordParameterInMySqlCli()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "mysql -u username -p\"password\"", matches[0])
}

func Test_HardcodedPasswordParameterInMySqlCli_3(t *testing.T) {
	testContent := `
mysql -p"password" -u username -h hostname database_name
`
	pattern := HardcodedPasswordParameterInMySqlCli()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "mysql -p\"password\"", matches[0])
}
