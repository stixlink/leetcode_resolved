package _004__median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r := find(nums1, nums2)
	return (float64(*r[0] + *r[1])) / 2
}

func find(nums1 []int, nums2 []int) (needValue [2]*int) {
	lenNums := len(nums1) + len(nums2)
	currentIndexResult := lenNums - 1
	var (
		needIndex1 int
		needIndex2 int
	)
	if lenNums%2 == 0 {
		needIndex1 = lenNums / 2
		needIndex2 = needIndex1 - 1
	} else {
		needIndex1 = lenNums / 2
		needIndex2 = needIndex1
	}

	i := len(nums1) - 1
	j := len(nums2) - 1
	for ; currentIndexResult >= -1; currentIndexResult-- {
		if needValue[0] != nil && needValue[1] != nil {
			return
		}
		if i < 0 {
			if needValue[0] == nil {
				r1 := j - needIndex1
				needValue[0] = &nums2[j-r1]
			}
			if needValue[1] == nil {
				r2 := j - needIndex2
				needValue[1] = &nums2[j-r2]
			}
			continue
		}

		if j >= 0 && nums1[i] < nums2[j] {
			if needIndex1 == currentIndexResult {
				needValue[0] = &nums2[j]
			}
			if needIndex2 == currentIndexResult {
				needValue[1] = &nums2[j]
			}
			j--
			continue
		}
		if j < 0 {
			if needValue[0] == nil {
				r1 := i - needIndex1
				needValue[0] = &nums1[i-r1]
			}
			if needValue[1] == nil {
				r2 := i - needIndex2
				needValue[1] = &nums1[i-r2]
			}
			return
		}
		if needIndex1 == currentIndexResult {
			needValue[0] = &nums1[i]
		}
		if needIndex2 == currentIndexResult {
			needValue[1] = &nums1[i]
		}
		i--
	}
	return
}
