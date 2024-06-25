package password_hashes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_BCrypt_positive(t *testing.T) {
	testContent := `
$2y$10$u.SFSH9FicoebQDcrd7Eeew9iDeMsbgUssDsF2/nxwsAI0ceCBXiC
$2y$10$oIUzEHU4tY/dWaOjhVqXKOVBMTFyKDXnXeoFV8xbYBrmJPv7U7lfe
$2y$10$nJC0eSD9dc4vFugUJ0rGpO7X2DdZSn364eHTghMPiDQdA1c0Y829i
$2y$10$QkdRayI7x89exQBmrQppoOiKW3n0eN1hUgDYmpYR7a/kGvQ6IDB3W
$2y$08$NYxEW8L5pWWO/4ZmQNAyPect0g5BViPspQzRbA0rODUEPG43uVBE2
$2y$12$8OluSUBf.04mo2cHyTfmbOEmOOb8qCxO47zqHAd.TWLmT5XjjtOeO
`
	pattern := BCrypt()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 6, len(matches))
	require.Equal(t, "$2y$10$u.SFSH9FicoebQDcrd7Eeew9iDeMsbgUssDsF2/nxwsAI0ceCBXiC", matches[0])
}

func Test_BCrypt_negative_TooLong(t *testing.T) {
	testContent := `
$2y$10$u.SFSH9FicoebQDcrd7Eeew9iDeMsbgUssDsF2/nxwsAI0ceCBXiC/TooLong
`
	pattern := BCrypt()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}

func Test_BCrypt_negative_TooShort(t *testing.T) {
	testContent := `
$2y$10$u.SFSH9FicoebQDcrd7Eeew9iDeMsbgUssDsF2/TooShort
`
	pattern := BCrypt()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}
