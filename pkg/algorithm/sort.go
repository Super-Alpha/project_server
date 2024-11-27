package algorithm

import (
	"fmt"
	"sort"
)

/*
冒泡排序, 时间复杂度: O(n^2) 空间复杂度: O(1)，稳定
算法步骤：

	1、前后两个数值进行大小比较，如果前面的数值比后面的大，则交换位置
	2、重复这个过程，直到所有元素都排好序。
*/
func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

/*
插入排序，时间复杂度: O(n^2)， 空间复杂度: O(1)， 稳定
算法步骤:

	1、将待排序的数组划分为已排序和未排序两部分，
	2、在未排序部分中取出一个元素，插入到已排序部分的合适位置，
	3、重复这个过程，直到所有元素都排好序。
*/
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		// 从后往前遍历，找到合适的位置插入
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

/*
选择排序，时间复杂度: O(n^2)， 空间复杂度: O(1)， 不稳定
算法步骤：

	1、首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	2、再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	3、重复第二步，直到所有元素均排序完毕。
*/
func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minValIndex := i
		// 从剩余元素中，寻找最小值的索引，并交换位置
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minValIndex] {
				minValIndex = j
			}
			nums[i], nums[minValIndex] = nums[minValIndex], nums[i]
		}
	}
}

/*
快速排序 时间复杂度: O(nlogn)， 空间复杂度: O(logn)， 不稳定

基本步骤：
1．先从数组中取出一个数作为基准数。（任意位置）
2．将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
3．再对左、右区间重复第二步，直到各区间只有一个数。
*/
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 基准值
	pivot := nums[left]
	i, j := left, right

	for i < j {
		// 寻找一个比基准数值大的数值
		for i < j && nums[i] < pivot {
			i++
		}
		// 寻找一个比基准数值小的数值
		for i < j && nums[j] > pivot {
			j--
		}
		// 找到之后，交换这两个数值
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i] = pivot

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

/*
归并排序：时间复杂度: O(nlogn)， 空间复杂度: O(n)， 稳定
(采用分治算法，然后对分治后的结果进行合并)
基本步骤：

	1、将n个元素分成个含n/2个元素的子序列。
	2、用合并排序法对两个子序列递归的排序。
	3、合并两个已排序的子序列已得到排序结果。
*/
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(leftNums, rightNums []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(leftNums) && j < len(rightNums) {
		if leftNums[i] < rightNums[j] {
			res = append(res, leftNums[i])
			i++
		} else {
			res = append(res, rightNums[j])
			j++
		}
	}
	res = append(res, leftNums[i:]...)
	res = append(res, rightNums[j:]...)
	return res
}

/*
堆排序：时间复杂度: O(nlogn)， 空间复杂度: O(1)， 不稳定
基本步骤：

	1、将待排序的数组转换为一个最大堆；
	2、依次取出堆顶元素放到已排序数组的末尾；
	3、重复这个过程直到所有的元素都排序好。
*/
func heapSort(nums []int) {
	length := len(nums)
	// 构建最大堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}
	// 依次将堆顶元素与最后一个元素交换，并重新调整堆
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

// 调整堆，使得以rootIndex为根节点的子树成为最大堆
func heapify(arr []int, n int, rootIndex int) {
	largest := rootIndex
	left, right := 2*rootIndex+1, 2*rootIndex+2
	// (若取小于号，则构建最小堆)
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	// (若取小于号，则构建最小堆)
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	// 如果最大值不是根节点，则交换根节点和最大值，并递归调整交换后的子树
	if largest != rootIndex {
		arr[rootIndex], arr[largest] = arr[largest], arr[rootIndex]
		heapify(arr, n, largest)
	}
}

func Sort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		// 小于是升序，大于是倒序
		return nums[i] < nums[j]
	})
}

func SortMain() {
	nums := []int{1, 3, 2, 0, 5, 4}
	//bubbleSort(nums)
	//insertionSort(nums)
	//selectionSort(nums)
	//quickSort(nums, 0, len(nums)-1)
	//fmt.Println(mergeSort(nums))
	//heapSort(nums)
	Sort(nums)
	fmt.Println(nums)
}
