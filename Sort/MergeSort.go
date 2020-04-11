package Sort

type MSort struct {
}

func (m MSort) Sort(arr []int) []int {

	if len(arr) > 1 {
		mid := int(len(arr) / 2)
		return merge(m.Sort(arr[:mid-1]), m.Sort(arr[mid:]))
	}

	return arr
}

func merge(arr1 []int, arr2 []int) []int {

	fArr := []int{}
	lArr1, lArr2 := len(arr1), len(arr2)
	for lArr1 > 0 || lArr2 > 0 {
		if lArr2 == 0 || lArr1 == 0 {
			if lArr2 == 0 {
				fArr = append(fArr, arr1[0])
			}
			if lArr1 == 0 {
				fArr = append(fArr, arr2[0])
			}

		}
	}

	return fArr
}
