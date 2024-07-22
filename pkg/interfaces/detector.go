package interfaces

type Detector interface {
	Check(OutputFormatter, LoadedFile) error
	Inject(llm LlmConnector)
}

type LlmConnector interface {
	GetBooleanResponse(string) (bool, error)
	GetResponseOutputModifier() string
}
