package sort

// MergeSort sorts a slice of integers
// Time Complexity: O(n log(n)), Memory: O(n)
// TODO: Sort slice in place instead of returning new slice (inconsistent with rest of sort API)
func MergeSort(vals []int) []int {
	if len(vals) <= 1 {
		return vals
	}
	pivot := len(vals) / 2
	left := MergeSort(vals[:pivot])
	right := MergeSort(vals[pivot:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var res = make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}
