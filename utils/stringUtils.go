package utils

import (
	"errors"
	"strings"
)

// ReplaceBefore Заменяет в строке [data] подстроку  до первого вхождения [separator] на [replacingData]
func ReplaceBefore(data, replacingData, separator string) (string, error) {
	var indexSeparator = strings.Index(data, separator)
	if indexSeparator == -1 {
		return "", errors.New("Не найден разделитель в строке: " + separator)
	}
	return strings.Replace(data, data[0:indexSeparator+1], replacingData, -1), nil
}
