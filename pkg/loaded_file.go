package pkg

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
)

type loadedFile struct {
	filename string
	filepath string
	content  []byte
}

func newLoadedFile(fileInfo interfaces.File, content []byte) interfaces.LoadedFile {
	return &loadedFile{
		filename: fileInfo.GetFilename(),
		filepath: fileInfo.GetFilepath(),
		content:  content,
	}
}

func (receiver loadedFile) GetFilename() string {
	return receiver.filename
}

func (receiver loadedFile) GetFilepath() string {
	return receiver.filepath
}

func (receiver loadedFile) GetContent() []byte {
	return receiver.content
}
