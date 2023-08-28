package leecode

func init() {
	Cmd(leecode{help: "两数之和",
		link: "https://leetcode.cn/problems/two-sum/",
		test: `
					[2,7,11,15] 9 [0,1]
					[3,2,4]     6 [1,2]
					[3,3]       6 [0,1]
		`,
		hand: func(nums []int, target int) []int {
			for i := 0; i < len(nums)-1; i++ {
				for j := i + 1; j < len(nums); j++ {
					if nums[i]+nums[j] == target {
						return []int{i, j}
					}
				}
			}
			return nil
		},
	})
}
