package interfaces

type Finding interface {
	GetName() string

	GetFile() LoadedFile

	GetContainsValue() bool
	GetValue() string

	GetIsCompleteFileImportant() bool

	GetFindingPriority() FindingPriority
}

type FindingInstance struct {
	File                    LoadedFile
	Name                    string
	Value                   string
	ContainsValue           bool
	IsCompleteFileImportant bool
	Priority                FindingPriority
}

func (f FindingInstance) GetFile() LoadedFile {
	return f.File
}

func (f FindingInstance) GetName() string {
	return f.Name
}

func (f FindingInstance) GetContainsValue() bool {
	return f.ContainsValue
}

func (f FindingInstance) GetValue() string {
	return f.Value
}

func (f FindingInstance) GetIsCompleteFileImportant() bool {
	return f.IsCompleteFileImportant
}

func (f FindingInstance) GetFindingPriority() FindingPriority {
	return f.Priority
}
