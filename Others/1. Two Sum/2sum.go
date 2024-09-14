package main

import "fmt"

func main() {
	ary := []int{1, 2, 3, 4, 5}
	ans := twoSum(ary, 6)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	traversedNums := make(map[int]int)

	for i, v := range nums {
		minused := target - v
		if m, exist := traversedNums[minused]; exist {
			return []int{m, i}
		}

		traversedNums[v] = i
	}

	return nil
}
