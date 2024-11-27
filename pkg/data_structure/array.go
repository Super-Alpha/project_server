package data_structure

import "fmt"

func Intersection(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)

	for i := 0; i < len(nums1); i++ {

		if func() bool {
			for j := 0; j < len(nums2); j++ {
				if nums1[i] == nums2[j] {
					return true
				}
			}
			return false
		}() {
			intersection = append(intersection, nums1[i])
		}
	}
	return intersection
}

func DeleteRepeatedElement(nums []int) int {
	temp := make([]int, 0)
	for _, v := range nums {
		if isExist(temp, v) {
			continue
		}
		temp = append(temp, v)
	}
	for k, val := range temp {
		nums[k] = val
	}
	fmt.Println(nums)
	return len(temp)
}

func isExist(num []int, target int) bool {
	for i := 0; i < len(num); i++ {
		if num[i] == target {
			return true
		}
	}
	return false
}

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func deleteSliceElement(nums []int, i int) {
	nums = append(nums[:i], nums[i+1:]...)
}

// nums := []int{1, 2, 3, 4, 5, 6, 7}
// rotate(nums, 9)
// fmt.Println(nums)
func rotate(nums []int, k int) {
	var j, x int
	for i := 0; i < k%len(nums); i++ {
		x = nums[len(nums)-1]
		copy(nums[j+1:], nums[j:len(nums)-j-1])
		nums[0] = x
		j = 0
	}
}

// s := "abc"
// t := "ahgdc"
// fmt.Println(isSubsequence(s, t))
func isSubsequence(s string, t string) bool {
	sList := []byte(s)
	tList := []byte(t)
	var j, count int
	for _, v := range sList {
		for k := j; k < len(tList); k++ {
			j++
			if v == tList[k] {
				count++
				break
			}
		}
	}
	return count == len(sList)
}

// [3,2,2,3]  3
// [2,2]
func removeElement(nums []int, val int) int {

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
		}
		if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			i++
			j--
		}
	}
	return i
}

// DeleteSliceElement 删除slice中的指定元素
func DeleteSliceElement(s []int, elem int) []int {
	j := 0
	for _, v := range s {
		if v != elem {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

// nums := []int{1, 2, 2, 3, 4, 4, 5, 6, 6}
// fmt.Println(removeDuplicates(nums))
func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
