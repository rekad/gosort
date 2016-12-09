package sort

// Lomuto partition scheme
// Time Complexity: O(n log(n)) average, O(n^2) worst case, Memory: O(log(n))
func QuickSort(vals []int) {
	if len(vals) <= 1 {
		return
	}

	// Using the Lomuto partition scheme:
	// Always pick last item as pivot
	var li = len(vals) - 1 // index of last element
	pivot := vals[li]
	i := 0 // where to put the pivot
	for j := 0; j < li; j++ {
		if vals[j] <= pivot {
			vals[i], vals[j] = vals[j], vals[i]
			i++
		}
	}
	vals[i], vals[li] = vals[li], vals[i]

	QuickSort(vals[:i])
	QuickSort(vals[i+1:])
}
