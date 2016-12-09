package sort

// RadixSort sorts a slice of integers
// TODO: Implement in place sorting
func RadixSort(vals []int) []int {

	largestNumber := getLargestNumber(vals)
	significantDigit := 1

	for largestNumber/significantDigit > 0 {
		buckets := [10][]int{}
		for i := 0; i < len(vals); i++ {
			bi := (vals[i] / significantDigit) % 10
			buckets[bi] = append(buckets[bi], vals[i])
		}
		vals = []int{}
		for i := 0; i < 10; i++ {
			vals = append(vals, buckets[i]...)
		}
		significantDigit *= 10
	}
	return vals
}

func getLargestNumber(vals []int) int {
	max := 0
	for i := 1; i < len(vals); i++ {
		if vals[i] >= vals[max] {
			max = i
		}
	}
	return vals[max]
}
