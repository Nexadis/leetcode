package twosum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := numMap[target-n]; ok {
			return []int{i, j}
		}
		numMap[n] = i
	}
	return nil
}
