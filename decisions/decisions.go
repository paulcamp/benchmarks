package decisions

// isInKnownSet is the original function
func isInKnownSet(classificationID int) bool {
	const (
		classification2  = 2
		classification4  = 4
		classification73 = 73
		classification85 = 85
		classification88 = 88
	)

	return classificationID == classification2 ||
		classificationID == classification4 ||
		classificationID == classification73 ||
		classificationID == classification85 ||
		classificationID == classification88
}

// Mapped version
var lookupMap = map[int]bool{
	2:  true,
	4:  true,
	73: true,
	85: true,
	88: true,
}

func isInMap(classificationID int) bool {
	return lookupMap[classificationID]
}

// Case version
func isInSwitchCase(classificationID int) bool {
	switch classificationID {
	case 2, 4, 73, 85, 88:
		return true
	}

	return false
}
