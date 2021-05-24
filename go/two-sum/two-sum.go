package two_sum

func twoSum(nums []int, target int) []int {
	tmp := map[int]int{}
	for i, num := range nums {
		mi := target - num //获取一个减法数字
		if _, ok := tmp[mi]; ok {
			return []int{tmp[mi], i}
		}
		tmp[mi] = i
	}
	for i, num := range nums {
		if _, ok := tmp[num]; ok {
			if tmp[num] != i {
				return []int{i, tmp[num]}
			}

		}
	}
	return nil
}
