package finetuning

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CheckProperCharacterFrequencyDistribution_FalsePositives(t *testing.T) {
	// if all upper case => probably placeholder
	require.Equal(t, false, CheckProperCharacterFrequencyDistribution("", "", "MY_PASSWORD"))
	require.Equal(t, false, CheckProperCharacterFrequencyDistribution("", "", "GHE_CREDENTIALS_PSW"))

	// * is too frequent to be a password
	require.Equal(t, false, CheckProperCharacterFrequencyDistribution("", "", "*******"))

	require.Equal(t, true, CheckProperCharacterFrequencyDistribution("", "", "ASDAJSKasd123456789"))
}

func Test_CheckProperCharacterFrequencyDistribution_TruePositives(t *testing.T) {
	require.Equal(t, true, CheckProperCharacterFrequencyDistribution("", "", "s2P7VSJc6O3dE2JijgEex"))
	require.Equal(t, true, CheckProperCharacterFrequencyDistribution("", "", "7Sn7sN"))
	require.Equal(t, true, CheckProperCharacterFrequencyDistribution("", "", "2g2tE88_$-VnYq5eyu:?JP}@H./Zk{:#TU"))
}

func Test_getCharacterStatistics_1(t *testing.T) {
	result := getCharacterStatistics("HelloWorld")
	require.Equal(t, 8, result.countOfLowerCaseCharacters)
	require.Equal(t, 2, result.countOfUpperCaseCharacters)
	require.Equal(t, 0, result.countOfSpecialCharacters)
	require.Equal(t, 10, result.totalNumberOfCharacters)
	require.Equal(t, 1, result.countPerCharacter['e'])
	require.Equal(t, 3, result.countPerCharacter['l'])
}
