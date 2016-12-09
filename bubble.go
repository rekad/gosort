package sort

// BubbleSort sorts a slice of integers
// Time Complexity: O(n^2), Memory: O(1)
func BubbleSort(vals []int) {
	for n := 0; n < len(vals); n++ {
		// If the sweep does not swap any numbers, the array is sorted
		swapped := false
		// After each sweep the last n numbers are in their final location
		for i := 0; i < len(vals)-1-n; i++ {
			if vals[i+1] < vals[i] {
				swapped = true
				vals[i], vals[i+1] = vals[i+1], vals[i]
			}
		}
		if !swapped {
			break
		}
	}
}
