package credential_assignments

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_fineTuningApprovesFinding_random_secrets(t *testing.T) {
	require.Equal(t, true, fineTuningApprovesFinding("", "password", "s2P7VSJc6O3dE2JijgEex"))
	require.Equal(t, true, fineTuningApprovesFinding("", "password", "7Sn7sN"))
	require.Equal(t, true, fineTuningApprovesFinding("", "password", "2g2tE88_$-VnYq5eyu:?JP}@H./Zk{:#TU"))
}

func Test_fineTuningApprovesFinding_placeholder(t *testing.T) {
	require.Equal(t, false, fineTuningApprovesFinding("", "password", "YOUR_PASSWORD"))
	require.Equal(t, false, fineTuningApprovesFinding("", "password", "API_KEY"))
	require.Equal(t, false, fineTuningApprovesFinding("", "password", "*********"))
	require.Equal(t, false, fineTuningApprovesFinding("", "password", "'GITHUB_PSW'"))
}
