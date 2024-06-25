package interfaces

type LoadedFile interface {
	GetFilename() string
	GetFilepath() string
	GetContent() []byte
}
