package Sort

type InsertionSort struct {
}

func (i InsertionSort) Sort(arr []int) []int {
	// 4,9,1,0,3
	if len(arr) >= 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		for i := 2; i < len(arr); i++ {
			index := i
			for j := i - 1; j >= 0; j-- {
				if arr[index] < arr[j] {
					arr[index], arr[j] = arr[j], arr[index]
					index = j
				}
			}

		}
	}
	return arr
}
