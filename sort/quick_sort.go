package sort

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
func quickSort(values []int, left, right int) {
	if left > right {
		return
	}
	mid := move(values, left, right)
	quickSort(values, left, mid-1)
	quickSort(values, mid+1, right)
}
func move(values []int, left, right int) int {
	v := values[left]
	i := left
	j := right
	for i < j {
		for values[j] >= v && i < j {
			j--
		}
		for values[i] <= v && i < j {
			i++
		}
		swap(values, i, j)
	}
	// 最终i==j
	swap(values, left, i)
	return i
}
func swap(values []int, left, right int) {
	tmp := values[left]
	values[left] = values[right]
	values[right] = tmp
}
