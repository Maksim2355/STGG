package utils

// RemoveFromIntSliceWithSaveOrder удаление элемента из слайса с сохранением порядка
func RemoveFromIntSliceWithSaveOrder(slice []int, indexOfItemBeingDeleted int) []int {
	return append(slice[:indexOfItemBeingDeleted], slice[indexOfItemBeingDeleted+1:]...)
}

// RemoveFromStringSliceWithSaveOrder удаление элемента из слайса с сохранением порядка
func RemoveFromStringSliceWithSaveOrder(slice []string, indexOfItemBeingDeleted int) []string {
	return append(slice[:indexOfItemBeingDeleted], slice[indexOfItemBeingDeleted+1:]...)
}

// RemoveFromStringMatrixWithSaveOrder удаление элемента из матрицы строк с сохранением порядка
func RemoveFromStringMatrixWithSaveOrder(slice [][]string, indexOfItemBeingDeleted int) [][]string {
	return append(slice[:indexOfItemBeingDeleted], slice[indexOfItemBeingDeleted+1:]...)
}

// RemoveFromIntSlice удаление числа из слайса без сохранения порядка элементов
func RemoveFromIntSlice(slice []int, indexOfItemBeingDeleted int) []int {
	slice[indexOfItemBeingDeleted] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// RemoveFromStringSlice удаление строки из слайса без сохранения порядка элементов
func RemoveFromStringSlice(slice []int, indexOfItemBeingDeleted int) []int {
	slice[indexOfItemBeingDeleted] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
