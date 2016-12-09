package sort

func InsertionSort(vals []int) {
	// iterate through array
	// if element is larger than the one before leave it
	// if it's bigger, insert it at the correct place and shift rest
	for i := 1; i < len(vals); i++ {
		for j := i - 1; j >= 0 && vals[j] > vals[j+1]; j-- {
			vals[j], vals[j+1] = vals[j+1], vals[j]
		}
	}
}
