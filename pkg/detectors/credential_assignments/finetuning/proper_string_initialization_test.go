package finetuning

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CheckProperStringInitialization(t *testing.T) {
	require.Equal(t, false, CheckProperStringInitialization(".go", "", "password"))
	require.Equal(t, false, CheckProperStringInitialization(".r", "", "password"))
	require.Equal(t, false, CheckProperStringInitialization(".php", "", "password"))
	require.Equal(t, false, CheckProperStringInitialization(".c", "", "password"))

	require.Equal(t, true, CheckProperStringInitialization(".go", "", "\"password\""))
	require.Equal(t, true, CheckProperStringInitialization(".r", "", "'password'"))
	require.Equal(t, true, CheckProperStringInitialization(".php", "", "'password'"))
	require.Equal(t, true, CheckProperStringInitialization(".c", "", "'password'"))
}
