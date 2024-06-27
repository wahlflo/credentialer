package credential_assignments

var interestingVariableNames = []string{
	"password",
	"secret",
	"token",
	"pwd",
	"pass",
	"passwd",
	//	"key",		<- generates too many false positives, ToDo: ignore this until Machine Learning model is integrated
	"пароль",   // Russian for password
	"passwort", // German for password
}
