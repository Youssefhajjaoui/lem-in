package utils

// this is used so i don't pass by refrence
func CopyMap(original map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}
