package utils

// MergeMaps Сливает две карты в одну
func MergeMaps(firstMap, secondMap map[interface{}]interface{}) map[interface{}]interface{} {
	for keySecondMap, valueSecondMap := range secondMap {
		firstMap[keySecondMap] = valueSecondMap
	}
	return firstMap
}
