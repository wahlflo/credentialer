package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_AwsAccessKeyId_positive_1(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID=ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_2(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID="ASIAIOSFODNN7EXAMPLE"`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_3(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID='ASIAIOSFODNN7EXAMPLE'`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_4(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID= ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_5(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID:ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_6(t *testing.T) {
	testContent := `AWS_ACCESS_KEY_ID:= ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
}

func Test_AwsAccessKeyId_positive_7(t *testing.T) {
	testContent := `
ASIAIOSFODNN7EXAMPLE -
ASIAIOSFODNN7EXAMPLE
`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[1])
}

func Test_AwsAccessKeyId_positive_8(t *testing.T) {
	testContent := `ASIAIOSFODNN7EXAMPLE -
ASIAIOSFODNN7EXAMPLE
`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[1])
}

func Test_AwsAccessKeyId_positive_9(t *testing.T) {
	testContent := `ASIAIOSFODNN7EXAMPLE - ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[1])
}

func Test_AwsAccessKeyId_positive_10(t *testing.T) {
	testContent := `ASIAIOSFODNN7EXAMPLE
ASIAIOSFODNN7EXAMPLE`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[0])
	require.Equal(t, "ASIAIOSFODNN7EXAMPLE", matches[1])
}

func Test_AwsAccessKeyId_negative_too_long(t *testing.T) {
	testContent := `
AWS_ACCESS_KEY_ID=ASIAIOSFODNN7EXAMPLETOOLONG
`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}

func Test_AwsAccessKeyId_negative_prefix(t *testing.T) {
	testContent := `
AWS_ACCESS_KEY_ID=asASIAIOSFODNN7EXAMPLE
`
	pattern := AwsAccessKeyId()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}

func Test_AwsSecretKey_positive(t *testing.T) {
	testContent := `
"123456789+123456789+123456789+123456789+"
'0000000000000000000000000000000000000000'
`
	pattern := AwsSecretKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
}

func Test_AwsSecretKey_negative(t *testing.T) {
	testContent := `
123456789+123456789+123456789+123456789+
0000000000000000000000000000000000000000
8d75008285babd33fd02e4160a6b1232d9521566
`
	pattern := AwsSecretKey()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}
