package Sort

//Worst-case performance	{\displaystyle O(n^{2})}O(n^{2}) comparisons, {\displaystyle O(n^{2})}O(n^{2}) swaps
//Best-case performance	{\displaystyle O(n)}O(n) comparisons, {\displaystyle O(1)}O(1) swaps
//Average performance	{\displaystyle O(n^{2})}O(n^{2}) comparisons, {\displaystyle O(n^{2})}O(n^{2}) swaps
//Worst-case space complexity

type BSort struct {
}

func (s BSort) Sort(arr []int) []int {
	swap, max := true, len(arr)

	for swap {
		iswap := false
		for i := 0; i < max-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				iswap = true
			}
		}
		swap = iswap
	}
	return arr
}

func (s BSort) SortByLength(arr []string) []string {
	swap, max := true, len(arr)

	for swap {
		iswap := false
		for i := 0; i < max-1; i++ {
			if len(arr[i]) > len(arr[i+1]) {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				iswap = true
			}
		}
		swap = iswap
	}
	return arr
}
