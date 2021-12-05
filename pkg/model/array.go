package model

func IsInclude(targetStrings []string, target string) bool {
	for _, element := range targetStrings {
		if element == target {
			return true
		}
	}
	
	return false
}
