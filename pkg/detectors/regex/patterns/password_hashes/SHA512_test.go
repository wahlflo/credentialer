package password_hashes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Sha512(t *testing.T) {
	testContent := `
linuxize:$6$zHvrJMa5Y690smbQ$z5zdL...:18009:0:120:7:14::
kalyani:$6$uUSXwCvO$Ic9kN9dS0BHN.NU.5h7rAcEQbtjPjqWpej5o5y7JlrQK0hdQrzKBZB1V6CowHhCpk25PaieLcJEqC6e02ExYA.:18917:0:99999:7:::
root:!:18613:0:99999:7:::
daemon:*:18613:0:99999:7:::
bin:*:18613:0:99999:7:::
systemd-resolve:*:18613:0:99999:7:::
systemd-timesync:*:18613:0:99999:7:::
messagebus:*:18613:0:99999:7:::
syslog:*:18613:0:99999:7:::
_apt:*:18613:0:99999:7:::
uuidd:*:18613:0:99999:7:::
tcpdump:*:18613:0:99999:7:::
remnux:$6$MqBJKwDlt.Dsx7Dp$aU6i2Mw.bvCRrS9DaRCo9k9ysoNbYc2WCwxvUmyCHX4dBs86SnAjscLKYE9Ld.fOXSd3WvIeougX53OjMK/DP1:18613:0:99999:7:::
`
	pattern := Sha512()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 2, len(matches))
	require.Equal(t, "$6$uUSXwCvO$Ic9kN9dS0BHN.NU.5h7rAcEQbtjPjqWpej5o5y7JlrQK0hdQrzKBZB1V6CowHhCpk25PaieLcJEqC6e02ExYA.", matches[0])
}
