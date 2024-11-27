package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"project_server/pkg/kafka"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
	"unsafe"
)

// return语句首先将变量放在临时变量之中，然后再返回数值；所以是先执行defer语句，然后再执行return语句；
func deferAndReturn() (a int) {
	a = 10
	defer func() {
		fmt.Println("start")
	}()

	defer func() {
		a += 2
		fmt.Println("demo", a)
	}()

	return a
}

func ParseUrl() {
	encoded := "company=%7B%7D&supplier=%7B%7D"

	// 解码 URL
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		log.Fatal(err)
	}

	// 解析参数
	values, err := url.ParseQuery(decoded)
	if err != nil {
		log.Fatal(err)
	}

	// 获取参数值
	company := values.Get("company")
	supplier := values.Get("supplier")

	fmt.Println("Decoded:", decoded)
	fmt.Println("Company:", company)
	fmt.Println("Supplier:", supplier)
}

type RetryFunc interface {
	Retry() (interface{}, error)
}

type Retry struct {
	Count        int
	Delay        time.Duration
	InternalFunc func() (interface{}, error)
}

func (r Retry) Retry() (interface{}, error) {
	var errs error
	for i := 0; i < 3; i++ {
		resp, err := r.InternalFunc()
		if err == nil {
			return resp, nil
		}
		errs = err
		time.Sleep(1 * time.Second)
	}
	return nil, errs
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	// 1、找到滑动窗口的头节点和尾节点
	tail := head
	count := k - 1
	for tail.Next != nil {
		tail = tail.Next
		count -= 1
		if count == 0 {
			reverseList(head).Next = tail.Next
			head = tail.Next
			count = k
		}
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := make([]int, 0)
					temp = append(temp, nums[i], nums[j], nums[k])
					if isCan(res, temp) {
						res = append(res, temp)
					}
					fmt.Println(res)
				}
			}
		}
	}
	return res
}

func isCan(list [][]int, nums []int) bool {
	for _, v := range list {
		if isSame(v, nums) {
			return false
		}
	}
	return true
}

func isSame(list1 []int, list2 []int) bool {
	m := make(map[int]int)
	for _, v := range list1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range list2 {
		if _, ok := m[v]; ok {
			m[v] -= 1
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func moveZeroes(nums []int) {
	j := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[j] == 0 {
			nums = append(nums[:j], nums[j+1:]...)
			nums = append(nums, 0)
		} else {
			j++
		}
	}
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if len(stack) > 1 {
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			} else if v == ']' && stack[len(stack)-1] != '[' {
				return false
			} else if v == '}' && stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, k)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if len(res) < k {
				res = append(res, []int{nums1[i], nums2[j]})
			} else {
				isSmallest(res, []int{nums1[i], nums2[j]})
			}
		}
	}

	return res
}

func isSmallest(nums [][]int, target []int) {
	for k, v := range nums {
		if v[0]+v[1] > target[0]+target[1] {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums, target)
		}
	}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	res := make([]int, 2)

	t1, t2 := true, true

	j := len(nums) - 1

	for i := 0; i <= j; i++ {
		if nums[i] == target && t1 {
			res[0] = i
			t1 = false
		}
		if nums[j] == target && t2 {
			res[1] = j
			t2 = false
		}
		if t2 {
			j--
		}
	}

	if t1 && t2 {
		return []int{-1, -1}
	}
	return res
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	temp := make([]int, 0)

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range m {
		temp = append(temp, v)
	}

	sort.Ints(temp)

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		for k1 := range m {
			if m[k1] == temp[len(temp)-1-i] {
				res = append(res, k1)
				delete(m, k1)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {

	mp := make(map[int]int)

	for i, v := range nums {
		temp := target - v
		j, ok := mp[temp]
		if ok {
			return []int{i, j}
		}
		mp[v] = i
	}

	return nil
}

func longestPrefix(str []string) string {
	if len(str) < 1 {
		return ""
	}
	p := str[0]
	for _, v := range str {
		for strings.Index(v, p) != 0 {
			if len(p) == 0 {
				return ""
			}
			p = p[:len(p)-1]
		}
	}
	return p
}

func removeNthNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	temp := head

	for temp != nil {
		count += 1
		temp = temp.Next
	}

	if count == n {
		return head.Next
	}

	temp = head

	for i := 0; i < count-n-1; i++ {
		temp = temp.Next
	}

	temp.Next = temp.Next.Next

	return head
}

func demo() {
	//threeSum([]int{-1, 0, 1, 2, -1, -4})
	//fmt.Println(isCan([][]int{{1, 2, 3}, {1, 2, 3}}, []int{1, 2, 3}))
	//fmt.Println(isSame([]int{1, 2, 3}, []int{1, 7, 3}))

	//nums := []int{0, 3, 1}
	//moveZeroes(nums)
	//fmt.Println(nums)
	//fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))

	// B+树的高度 H = logN/logB (其中N=数据量，B=每个节点能存储的索引数)
	// 假设N=10万量级，B=32（每个节点可以存储32个索引）
	// 那么 H = log2^17 / log2^6 = 17/6 即 2~3 之间

	//data_structure.Main()
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	//fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	//algorithm.SortMain()
	//algorithm.SearchMain()
}

func strToInt(str string) int {
	if str == "" {
		return 0
	}
	num := 0
	for i := 0; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		num = num*10 + int(str[i]-'0')
	}
	return num
}

func demos() {
	// 假设我们有一个非常大的int型slice
	numbers := generateSlice(1000000) // 生成一个包含100万个随机整数的切片
	// 要查找的目标数字
	target := 42
	// 创建一个共享的上下文，以便在找到目标时取消所有协程
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 创建一个等待组，以便等待所有协程完成
	var wg sync.WaitGroup
	// 设定并发协程的数量
	numRoutines := 10
	// 将大切片分割为小块给各个协程处理
	chunkSize := len(numbers) / numRoutines
	// 添加足够的等待项，以便等待所有协程完成
	wg.Add(numRoutines)
	// 启动协程进行搜索
	for i := 0; i < numRoutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numRoutines-1 {
			end = len(numbers) // 确保最后一个协程处理剩余的所有元素
		}
		go func(subset []int, ctx context.Context) {
			defer wg.Done() // 协程完成时通知等待组
			for _, num := range subset {
				select {
				case <-ctx.Done(): // 如果上下文被取消，退出循环
					return
				default:
					if num == target {
						fmt.Printf("Target found by routine: %d\n", num)
						cancel() // 取消上下文，通知其他协程停止搜索
						return
					}
				}
			}
		}(numbers[start:end], ctx)
	}
	// 等待所有协程完成
	wg.Wait()
	fmt.Println("Search completed.")
}

// 生成一个包含n个随机整数的切片
func generateSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(1000) // 假设数字范围是0-999
	}
	return numbers
}

func deleteElement(nums []int, i int) {
	nums = append(nums[:i], nums[i:]...)
}

// nums = [0,0,1,1,1,1,2,3,3]
// 7, nums = [0,0,1,1,2,3,3]
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	i := 2

	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6
func maxSubArray(nums []int) int {
	res := math.MinInt64
	sumValue := 0
	for i := 0; i < len(nums); i++ {
		if sumValue <= 0 {
			sumValue = nums[i]
		} else {
			sumValue += nums[i]
		}
		if sumValue > res {
			res = sumValue
		}
	}
	return res
}

// 最长连续序列长度
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	count := 1

	mp := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		if nums[i+1]-nums[i] == 1 {
			count += 1
		} else {
			mp[count] = 0
			count = 1
		}
	}
	mp[count] = 0

	res := 0

	for k := range mp {
		if k > res {
			res = k
		}
	}

	return res
}

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
end:
	for i = 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i+1 > len(strs[j]) {
				break end
			}
			if strs[j][i] != strs[0][i] {
				break end
			}
		}
	}

	return strs[0][:i]
}

func deleteSameElement(nums []int) []int {
	res := make([]int, 0)
	mp := make(map[int]int)
	for _, v := range nums {
		if _, ok := mp[v]; !ok {
			mp[v] = 0
			res = append(res, v)
		}
	}
	return res
}

// target = 7, nums = [2,3,1,2,4,3]
func minSubArrayLen(target int, nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}

	if s < target {
		return 0
	}

	minLength := len(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return 1
		}
		sumVal := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if sumVal < target {
				sumVal += nums[j]
			}
			if sumVal > target {
				break
			}
			if sumVal == target {
				if j-i+1 < minLength {
					minLength = j - i + 1
				}
				break
			}
		}
	}
	return minLength
}

// 输入：s = "the sky is blue"
// 输出："blue is sky the"
func reverseWords(s string) string {

	strs := strings.Split(strings.Trim(s, " "), " ")

	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}

	res := ""

	for k, v := range strs {
		if v != "" && k != len(strs)-1 {
			res += v + " "
		}
		if v != "" && k == len(strs)-1 {
			res += v
		}
	}

	return res
}

// 输入：haystack = "sadbutsad", needle = "sad"
// 输出：0
func strStr(haystack string, needle string) int {

	if haystack == needle {
		return 0
	}

	length := len(needle)

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+length <= len(haystack) && haystack[i:i+length] == needle {
				return i
			}
		}
	}

	return -1
}

// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {

	mp := make(map[string][]string)

	for _, v := range strs {

		key := sortString(v)

		if _, ok := mp[key]; ok {
			mp[key] = append(mp[key], v)
		} else {
			mp[key] = []string{v}
		}
	}

	res := make([][]string, 0)

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

func sortString(str string) string {

	s := []byte(str)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return string(s)
}

func sortList(head *ListNode) *ListNode {

	return nil
}

//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引2

func findPeakElement(nums []int) int {

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
	}

	if nums[len(nums)-1] > nums[0] && nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	return 0
}

// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
// [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
func kSmallestPairsTest(nums1 []int, nums2 []int, k int) [][]int {
	res := make(map[int][][]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if _, ok := res[nums1[i]+nums2[j]]; ok {
				res[nums1[i]+nums2[j]] = append(res[nums1[i]+nums2[j]], []int{nums1[i], nums2[j]})
			} else {
				res[nums1[i]+nums2[j]] = [][]int{{nums1[i], nums2[j]}}
			}
		}
	}

	result := make([][]int, 0)
	for _, v := range res {
		result = append(result, v...)
	}

	return result[:k]
}

func compoundInterest() {
	principal, rate, years := 20000.0, 0.02, 1.0

	amount := principal * math.Pow(1+rate, years)

	interest := amount - principal

	fmt.Printf("TotalAmount = %.2f\nPrincipal = %.2f\nInterest = %.2f\n", amount, principal, interest)
}

// "()[]{}"  ([}}])
func isValidString(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			switch v {
			case ')':
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)

	for i := k - 1; i < len(nums); i++ {

		if len(res) != 0 {
			if res[len(res)-1] > nums[i] && isContain(nums[i-k+1:i+1], res[len(res)-1]) {
				res = append(res, res[len(res)-1])
			} else {
				res = append(res, maxValue(nums[i-k+1:i+1]))
			}
		} else {
			res = append(res, maxValue(nums[i-k+1:i+1]))
		}
	}
	return res
}
func maxValue(nums []int) int {
	val := nums[0]
	for _, v := range nums {
		if v > val {
			val = v
		}
	}
	return val
}

func isContain(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

type Scalar interface {
	int | float64 | float32 | string
}

func Plus[T Scalar](a, b T) T {

	return a + b
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, right, target)
	}
}

//给定一个字符串 s 请你找出其中不含有重复字符的最长子串的长度。
//
//示例 1:
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func lengthOfLongestSubstring(s string) int {
	res := 0
	for left, v := range []byte(s) {
		maxVal := 1
		table := make([]byte, 0)
		table = append(table, v)
		for right := left + 1; right < len(s); right++ {
			if isExists(table, s[right]) {
				break
			} else {
				table = append(table, s[right])
				maxVal += 1
			}
		}
		if maxVal > res {
			res = maxVal
		}
	}

	return res
}

func isExists(s []byte, target byte) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// MbSubStr fmt.Println(MbSubStr("世界2024年9月11日", 2, 5))
func MbSubStr(str string, offset, length int) string {
	strLen := utf8.RuneCountInString(str)
	if offset >= strLen {
		return ""
	}

	end := offset + length
	if end > strLen {
		end = strLen
	}

	s := []rune(str)

	return string(s[offset:end])
}

// nums := []int{1, 10, 8, 5, 6, 11}
// fmt.Println(handleNums(nums))
func handleNums(nums []int) []int {
	s := make([]int, 0)
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = -1
	}

	for i, num := range nums {
		for len(s) > 0 && nums[s[len(s)-1]] < num {
			index := s[len(s)-1]
			s = s[:len(s)-1]
			res[index] = num
		}
		s = append(s, i)
	}

	return res
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度
func lengthOfLongestSubstrings(s string) int {
	mp := make(map[byte]int)

	n := len(s)

	r, a := -1, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(mp, s[i-1])
		}

		for r+1 < n && mp[s[r+1]] == 0 {
			mp[s[r+1]] += 1
			r += 1
		}

		a = maxVal(a, r-i+1)
	}

	return a
}

func maxVal(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// [3,2,-1,3,-2]

// [3,2,-1,3]
func maxNum(nums []int) int {

	maxSum := nums[0]

	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func compoundInterests(principal float64, rate float64, timesCompounded int, years int) float64 {
	return principal * math.Pow(1+rate/float64(timesCompounded), float64(timesCompounded*years))
}

func main() {
	//kafka.Producer()

	kafka.Consumer()

	//kafka.ConsumerGroupMain()

	//kafka.Broker()

	//kafka.Main()
}
