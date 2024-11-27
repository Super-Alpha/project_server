package algorithm

import "fmt"

// 滑动窗口求最大值
// nums = [1,3,-1,-3,5,3,6,7], 和 k = 3 输出: [3,3,5,5,6,7]

func MaxSlidingWindow(nums []int, k int) []int {
	if k == 1 || len(nums) == 0 {
		return nums
	}

	res := make([]int, 0)
	// 单向递减队列（维护滑动窗口最大值（即队首元素））
	deque := make([]int, 0)

	for i := range nums {
		// 队内小于当前元素的数值出队
		for len(deque) > 0 && nums[i] > deque[len(deque)-1] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, nums[i])

		// 滑动窗口时，并且队首元素不在滑动窗口范围内，则队首元素出队
		if i >= k && len(deque) > 0 && nums[i-k] == deque[0] {
			deque = deque[1:]
		}

		// 滑动窗口时，将队列首元素加入结果数组
		if i >= k-1 && len(deque) > 0 {
			res = append(res, deque[0])
		}
	}
	return res
}

func main() {
	res := MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(res)
}
