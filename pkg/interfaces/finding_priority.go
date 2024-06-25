package interfaces

import "errors"

type FindingPriority int

const (
	FindingPriorityInformative FindingPriority = iota + 1
	FindingPriorityLow
	FindingPriorityMedium
	FindingPriorityHigh
)

func (p FindingPriority) ToString() string {
	switch p {
	case FindingPriorityInformative:
		return "informative"
	case FindingPriorityLow:
		return "low"
	case FindingPriorityMedium:
		return "medium"
	case FindingPriorityHigh:
		return "high"
	}
	panic(errors.New("given priority is not supported"))
}
