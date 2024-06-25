package private_keys

import (
	"github.com/stretchr/testify/require"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"testing"
)

func createTestFinding(match string) interfaces.FindingInstance {
	return interfaces.FindingInstance{
		Value:    match,
		Priority: interfaces.FindingPriorityHigh,
	}
}

func Test_private_key_pkcs1_rsa(t *testing.T) {
	testContent := `
-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf9Cnzj4p4WGeKLs1Pt8Qu
KUpRKfFLfRYC9AIKjbJTWit+CqvjWYzvQwECAwEAAQJAIJLixBy2qpFoS4DSmoEm
o3qGy0t6z09AIJtH+5OeRV1be+N4cDYJKffGzDa88vQENZiRm0GRq6a+HPGQMd2k
TQIhAKMSvzIBnni7ot/OSie2TmJLY4SwTQAevXysE2RbFDYdAiEBCUEaRQnMnbp7
9mxDXDf6AU0cN/RPBjb9qSHDcWZHGzUCIG2Es59z8ugGrDY+pxLQnwfotadxd+Uy
v/Ow5T0q5gIJAiEAyS4RaI9YG8EWx/2w0T67ZUVAw8eOMB6BIUg0Xcu+3okCIBOs
/5OiPgoTdSy7bcF9IGpSE8ZgGKzgYQVZeN97YE00
-----END RSA PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityHigh)
}

func Test_private_key_pkcs1_rsa_encrypted(t *testing.T) {
	testContent := `
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,5C724CE55C702828F3F74B555F594366

odKAmV6AbsoWsyL3thUoYVDEJAsQl8RrH+JuQ9HWUnDLunDdLEM6oNl15XP1xLOH
z3bEq1rvATiQmAByKNOiVujd1gsq7JxfQYDdHRzDhZZrUstnetvGTDBtMHmhzbBX
Oih+1q3eA2RMQ5izXOEkyMKrWWlcKMWVJzMSYjFeFJB8D8wJNmq1ArNCO3uXfwkZ
uMnMhYhx/OYvCs4sMWKe5/etyR2gz0Fvp6VDUa0jNRvoad+8/pHK7KDxB8nW5Kgm
pSjfkl1Ut3zChtwEuAFnSDuypbrODBdphZHD40WmX0f69VKKs44vsKCHr8nzJ8R5
dw+2Ggyq5W5hl3PDTMTqn8Pc+cwmPdVe4bkNqxbCHe2omZXpNIgC31wrMBvkyUYv
pY8rMoBXqgm9hC5JsXzn6Z6X1kpGFhDjkNSdzx4jYzw=
-----END RSA PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityInformative)
}

func Test_private_key_pkcs8(t *testing.T) {
	testContent := `
-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAqPfgaTEWEP3S9w0t
gsicURfo+nLW09/0KfOPinhYZ4ouzU+3xC4pSlEp8Ut9FgL0AgqNslNaK34Kq+NZ
jO9DAQIDAQABAkAgkuLEHLaqkWhLgNKagSajeobLS3rPT0Agm0f7k55FXVt743hw
Ngkp98bMNrzy9AQ1mJGbQZGrpr4c8ZAx3aRNAiEAoxK/MgGeeLui385KJ7ZOYktj
hLBNAB69fKwTZFsUNh0CIQEJQRpFCcydunv2bENcN/oBTRw39E8GNv2pIcNxZkcb
NQIgbYSzn3Py6AasNj6nEtCfB+i1p3F35TK/87DlPSrmAgkCIQDJLhFoj1gbwRbH
/bDRPrtlRUDDx44wHoEhSDRdy77eiQIgE6z/k6I+ChN1LLttwX0galITxmAYrOBh
BVl433tgTTQ=
-----END PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityHigh)
}

func Test_private_key_pkcs8_encrypted(t *testing.T) {
	testContent := `
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIIBvTBXBgkqhkiG9w0BBQ0wSjApBgkqhkiG9w0BBQwwHAQIpZHwLtkYRb4CAggA
MAwGCCqGSIb3DQIJBQAwHQYJYIZIAWUDBAECBBCCGsoP7F4bd8O5I1poTn8PBIIB
YBtM1tgqsAQgbSZT0475aHufzFuJuPWOYqiHag8OUKMeZuxVHndElipEY2V5lS9m
wddwtWaGuYD/Swcdt0Xht8U8BF0SjSyzQ4YtRsG9CmEHYhWmQ5AqK1W3mDUApO38
Cm5L1HrHV4YJnYmmK9jgq+iWlLFDmB8s4TA6kMPWbCENlpr1kEXz4hLwY3ylH8XW
I65WX2jGSn61jayCwpf1HPFBPDUaS5s3f92aKjk0AE8htsDBBiCVS3Yjq4QSbhfz
uNIZ1TooXT9Xn+EJC0yjVnlTHZMfqrcA3OmVSi4kftugjAax4Z2qDqO+onkgeJAw
P75scMcwH0SQUdrNrejgfIzJFWzcH9xWwKhOT9s9hLx2OfPlMtDDSJVRspqwwQrF
QwinX0cR9Hx84rSMrFndxZi52o9EOLJ7cithncoW1KOAf7lIJIUzP0oIKkskAndQ
o2UiZsxgoMYuq02T07DOknc=
-----END ENCRYPTED PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityInformative)
}

func Test_private_key_openssh(t *testing.T) {
	testContent := `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAVwAAAAdzc2gtcn
NhAAAAAwEAAQAAAEEAqPfgaTEWEP3S9w0tgsicURfo+nLW09/0KfOPinhYZ4ouzU+3xC4p
SlEp8Ut9FgL0AgqNslNaK34Kq+NZjO9DAQAAATB+9/CSfvfwkgAAAAdzc2gtcnNhAAAAQQ
Co9+BpMRYQ/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2y
U1orfgqr41mM70MBAAAAAwEAAQAAAEAgkuLEHLaqkWhLgNKagSajeobLS3rPT0Agm0f7k5
5FXVt743hwNgkp98bMNrzy9AQ1mJGbQZGrpr4c8ZAx3aRNAAAAIBOs/5OiPgoTdSy7bcF9
IGpSE8ZgGKzgYQVZeN97YE00AAAAIQCjEr8yAZ54u6Lfzkontk5iS2OEsE0AHr18rBNkWx
Q2HQAAACEBCUEaRQnMnbp79mxDXDf6AU0cN/RPBjb9qSHDcWZHGzUAAAAXcGhwc2VjbGli
LWdlbmVyYXRlZC1rZXkBAgME
-----END OPENSSH PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	require.Equal(t, privateKeyQualityCheck(createTestFinding(testContent)).GetFindingPriority(), interfaces.FindingPriorityHigh)

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityHigh)
}

func Test_private_key_openssh_encrypted(t *testing.T) {
	testContent := `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABCmZ5U5Eu
qcHFCIfF9gfNrvAAAAEAAAAAEAAABXAAAAB3NzaC1yc2EAAAADAQABAAAAQQCo9+BpMRYQ
/dL3DS2CyJxRF+j6ctbT3/Qp84+KeFhnii7NT7fELilKUSnxS30WAvQCCo2yU1orfgqr41
mM70MBAAABMM5HiDWh0Vf5BWUVoso9wTFYoNtxPBfHa3NQk+i/1XLx0ZQbYfurzUkE54Zi
gVPaGYMHbK1whuxSmRD6JlI3w+lENdjgTX/PR6netDsYKO0AbFxKEmzAItGbJAekcqdELA
QjEvPDO6BQaBKrBNqrj9S21uA7pNZyqV6uZL7pSZR89B8OmLpN5v5IzXFkjmYzwpT71b+C
gZ0q2mOH60b+1h1mN3jFjLPVIrpUiUzDhscX6hjd1XD3a69CjsN5IKUbEVp0zb4QNCz7RU
a4fY8yTCwSQ3VBloX1+ysKv/OX75Bb4WtLpUz3V/gehiYuY9Qm4Cq9wfXI3WgBqFld/8z+
qmrujXsdNGHAGaHCLD5TQLOn3ZBpEzfLBLcOka89zUAs+JDfHOB/UJ3T1raVNriM3Gc=
-----END OPENSSH PRIVATE KEY-----
`
	pattern := PrivateKeyGeneric()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))

	privateKeysEnsureProperPriority(t, testContent, interfaces.FindingPriorityInformative)
}

func privateKeysEnsureProperPriority(t *testing.T, testData string, priority interfaces.FindingPriority) {
	finding := interfaces.FindingInstance{
		Value:    testData,
		Priority: interfaces.FindingPriorityHigh,
	}

	require.Equal(t, priority, privateKeyQualityCheck(finding).GetFindingPriority())
}
