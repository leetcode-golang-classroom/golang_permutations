package sol

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := [][]int{}
	used := make([]bool, len(nums))
	var genPerm func(start int, perm []int)
	genPerm = func(start int, perm []int) {
		if start == len(nums) { // reach the end
			temp := make([]int, len(perm))
			copy(temp, perm)
			result = append(result, temp)
			return
		}
		// from start find not used element to permutation
		for idx := 0; idx < len(nums); idx++ {
			if !used[idx] { // not used index
				used[idx] = true
				// take nums[idx]
				perm = append(perm, nums[idx])
				// try next
				genPerm(start+1, perm)
				// rollback
				perm = perm[:len(perm)-1]
				used[idx] = false
			}
		}
	}
	genPerm(0, []int{})
	return result
}
