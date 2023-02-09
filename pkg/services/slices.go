package services

func ContainSlice(slice []string, target string) bool {
	for _, el := range slice {
		if el == target {
			return true
		}
	}
	return false
}
