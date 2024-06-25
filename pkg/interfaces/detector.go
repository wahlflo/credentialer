package interfaces

type Detector interface {
	Check(OutputFormatter, LoadedFile) error
}
