package interfaces

type OutputFormatter interface {
	Start()
	AddFinding(finding Finding)
	Finished()
}
