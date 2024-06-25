package credential_assignments

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getFindingsForVariableName_Case01(t *testing.T) {
	testContent := `
"https://maps.googleapis.com/maps/api/js?key=YOUR_API_KEY&callback=initMap"
password=ABC121
password = ABC122
password ="ABC123"
password="ABC124"
password = "ABC125"
password ='ABC126'
password = 'ABC127'
password: ABC128
password:= "ABC129"
let mut password = String::from("Hallo");		// is picked up by two different expressions
`
	matches := getFindings("", []byte(testContent))

	for _, match := range matches {
		fmt.Println(match.usedRegex)
		fmt.Println(match.value)
		fmt.Println("######################")
	}

	require.Equal(t, 12, len(matches))
}

func Test_getFindings_Case02_Rust(t *testing.T) {
	testContent := `
let mut password = String::from("Hallo");
`
	matches := getFindings(".rs", []byte(testContent))

	for _, match := range matches {
		fmt.Println(match.usedRegex)
		fmt.Println(match.value)
		fmt.Println("######################")
	}

	require.Equal(t, 1, len(matches))
}
func Test_getFindings_Case03(t *testing.T) {
	testContent := `
AWS_SECRET_KEY = "asdasd";
`
	matches := getFindings(".py", []byte(testContent))

	for _, match := range matches {
		fmt.Println(match.usedRegex)
		fmt.Println(match.variableName)
		fmt.Println(match.value)
		fmt.Println("######################")
	}

	require.Equal(t, 1, len(matches))
}

func Test_getFindings_negative_1(t *testing.T) {
	testContent := `
your password is ABC123
token = $o->{clock};
`
	matches := getFindings("", []byte(testContent))

	require.Equal(t, 0, len(matches))
}

func Test_getFindings_negative_2(t *testing.T) {
	testContent := `
GitHubToken: {{GITHUB_TOKEN}}
`
	matches := getFindings("", []byte(testContent))

	require.Equal(t, 0, len(matches))
}
