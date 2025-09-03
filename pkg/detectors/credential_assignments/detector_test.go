package credential_assignments

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getFindings_Case01(t *testing.T) {
	testContent := `
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
api_key = "key01"
apikey = 'key02'
api key: key03
api-key := "key04"
let mut apikey = String::from("key05");
API_KEY = "KEY06"
`
	matches := getFindings("", []byte(testContent), nil)

	for _, match := range matches {
		fmt.Println(match.usedRegex)
		fmt.Println(match.value)
		fmt.Println("######################")
	}

	require.Equal(t, 18, len(matches))
}

func Test_getFindings_Case02_Rust(t *testing.T) {
	testContent := `
let mut password = String::from("Hallo");
`
	matches := getFindings(".rs", []byte(testContent), nil)

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
	matches := getFindings(".py", []byte(testContent), nil)

	for _, match := range matches {
		fmt.Println(match.usedRegex)
		fmt.Println(match.variableName)
		fmt.Println(match.value)
		fmt.Println("######################")
	}

	require.Equal(t, 1, len(matches))
}

func Test_getFindings_Case04(t *testing.T) {
	testContent := `
password = ABC123
`
	matches := getFindings("", []byte(testContent), nil)

	require.Equal(t, 1, len(matches))
	require.Equal(t, "password", matches[0].variableName)
	require.Equal(t, "ABC123", matches[0].value)
}

func Test_getFindings_negative_1(t *testing.T) {
	testContent := `
your password is ABC123
token = $o->{clock};
`
	matches := getFindings(".ps1", []byte(testContent), nil)

	require.Equal(t, 0, len(matches))
}

func Test_getFindings_negative_placeholder(t *testing.T) {
	testContent := `
"https://maps.googleapis.com/maps/api/js?Apikey=YOUR_API_KEY&callback=initMap"
`
	matches := getFindings("", []byte(testContent), nil)
	require.Equal(t, 0, len(matches))
}

func Test_Finetuning_FalsePositives(t *testing.T) {
	ensureFindingIsNotApproved(t, "GitHubToken", "{{GITHUB_TOKEN}}")
	ensureFindingIsNotApproved(t, "Password", "\"MY_PASSWORD\"")
	ensureFindingIsNotApproved(t, "Password", "\"your_secret\"")
	ensureFindingIsNotApproved(t, "keyword", "(value:'STDLOGR'))")
	ensureFindingIsNotApproved(t, "key", "key:term,negate:!f,type:custom")
	ensureFindingIsNotApproved(t, "password", "System.getenv(\\\"")
}

func Test_Finetuning_TruePositives(t *testing.T) {
	// random generated passwords
	ensureFindingIsApproved(t, "password", "&GIw681L@7nv")
	ensureFindingIsApproved(t, "password", "NIPAQiIgvThakkNwDL")
	ensureFindingIsApproved(t, "password", "a063x8fdj37m5hmf91")
	ensureFindingIsApproved(t, "password", "87b73`=n0<")
	ensureFindingIsApproved(t, "password", "ba7t6260v2")
	ensureFindingIsApproved(t, "password", "lL059[sz")
}

func ensureFindingIsApproved(t *testing.T, variableName string, password string) {
	require.Equal(t, true, fineTuningApprovesFinding("", variableName, password))
}

func ensureFindingIsNotApproved(t *testing.T, variableName string, password string) {
	require.Equal(t, false, fineTuningApprovesFinding("", variableName, password))
}
