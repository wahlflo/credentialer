package llms

type LlmConnector interface {
	CheckConnection() error
	GetBooleanResponse(string) (bool, error)
	GetResponseOutputModifier() string
}
