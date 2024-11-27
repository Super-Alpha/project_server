package algorithm

import "fmt"

// 二分查找(针对有序数组)，时间复杂度为O(logn)，空间复杂度为O(1)
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2 //重要‼️
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 插值查找(针对有序数组)，平均时间复杂度为O(log logn)，最差为O(logn)，空间复杂度为O(1)
func searchInsert(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		// 计算插值位置
		pos := low + (target-arr[low])*(high-low)/(arr[high]-arr[low])

		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	// 未找到目标元素
	return -1
}

func SearchMain() {
	arr := []int{1, 2, 5, 7, 9}
	target := 5
	index := searchInsert(arr, target)
	fmt.Printf("目标值 %d 在数组中的索引为：%d", target, index)
}
