package _033_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if target < nums[0] {
		for j := len(nums) - 1; j > 0; j-- {
			if target == nums[j] {
				return j
			}
		}
	} else {
		for g := 0; g < len(nums); g++ {
			if target == nums[g] {
				return g
			}
		}
	}

	return -1

}
