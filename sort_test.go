package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type testCase struct {
	input  []int
	output []int
}

const N = 50000

var bigSlice = make([]int, N)
var bigSliceSorted = make([]int, N)

func init() {
	for i := 0; i < N; i++ {
		bigSlice[i] = rand.Int()
	}
	copy(bigSliceSorted, bigSlice)
	sort.Ints(bigSliceSorted)
}

func generateTestCases() []testCase {
	var bigSliceCopy = make([]int, N)
	copy(bigSliceCopy, bigSlice)
	return []testCase{
		{[]int{2, 0, 1, 9, 3, 8, 4, 7, 6, 5}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{bigSliceCopy, bigSliceSorted},
	}
}

func TestBubbleSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if BubbleSort(test.input); !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Not sorted correctly: %v", test.input)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if SelectionSort(test.input); !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Not sorted correctly: %v", test.input)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if got := MergeSort(test.input); !reflect.DeepEqual(got, test.output) {
			t.Errorf("Not sorted correctly: %v", test.input)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if QuickSort(test.input); !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Not sorted correctly: %v", test.input)
		}
	}
}

func TestRadixSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if got := RadixSort(test.input); !reflect.DeepEqual(got, test.output) {
			t.Errorf("Not sorted correctly: %v", got)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range generateTestCases() {
		if InsertionSort(test.input); !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Not sorted correctly: %v", test.input)
		}
	}
}
