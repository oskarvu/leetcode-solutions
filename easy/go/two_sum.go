func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for index, num := range nums {
		if pairIndex, ok := numsMap[target-num] ; ok {
			return []int{index, pairIndex}
		}
		numsMap[num] = index
	}
	return nil
}