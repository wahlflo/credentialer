package interfaces

func GetFindingWithHighestPriority(findings []Finding) Finding {
	selectedItem := findings[0]
	for _, i := range findings[1:] {
		if i.GetFindingPriority() > selectedItem.GetFindingPriority() {
			selectedItem = i
		}
	}
	return selectedItem
}
