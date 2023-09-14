package main

func twoSum(nums []int, target int) []int {
	mapp := make(map[int]int)

	arr := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		value, containsKey := mapp[target-nums[i]]
		if containsKey {
			arr[0] = value
			arr[1] = i
		} else {
			mapp[nums[i]] = i
		}
	}

	return arr
}
