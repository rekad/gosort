package sort

const MaxInt = int(^uint(0) >> 1)

// SelectionSort sorts a slice of integers
// Time Complexity: O(n^2), Memory: O(1)
func SelectionSort(vals []int) {
	for n := 0; n < len(vals)-1; n++ {
		var min = n
		for i := n + 1; i < len(vals); i++ {
			if vals[i] < vals[min] {
				min = i
			}
		}
		vals[n], vals[min] = vals[min], vals[n]
	}
}
