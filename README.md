# golang_permutations

Given an array `nums` of distinct integers, return *all the possible permutations*. You can return the answer in **any order**.

## Examples

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

```

**Example 2:**

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]

```

**Example 3:**

```
Input: nums = [1]
Output: [[1]]

```

**Constraints:**

- `1 <= nums.length <= 6`
- `10 <= nums[i] <= 10`
- All the integers of `nums` are **unique**.

## 解析

給一個整數陣列 nums

要求寫一個演算法找出所有排列的方法

可以考慮使用決策樹

![](https://i.imgur.com/iZOZVs2.png)

透過決策樹每次選擇一個陣列元素當作起點

並且紀錄已經使用過的元素

然後逐步帶入尚未使用的元素

## 程式碼
```go
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

```
## 困難點

1. 理解如何透過窮舉的方式來找出所有可能的組合
2. 理解如何循序窮舉

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity