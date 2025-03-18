package algorithm

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func insert(head *Node, index int, data int) *Node {
	p := head
	for i := 0; i < index && p != nil; i++ {
		p = p.next
	}
	if p == nil {
		return nil
	}

	newNode := &Node{data: data}
	newNode.next = p.next
	p.next = newNode

	return head
}

func length(head *Node) int {
	if head == nil {
		return 0
	}
	size := 0
	p := head
	for p != nil {
		size += 1
		p = p.next
	}
	return size
}

func printElement(head *Node) {
	p := head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

// 反转链表
func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var prev *Node // 前一节点
	cur := head    // 当前节点
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 将两个有序链表合并为一个有序链表
// 1 -> 2 -> 2 -> 3 -> 4
// 1 -> 3 -> 5
func mergeTwoLists(node1, node2 *Node) *Node {
	root := &Node{} // 固定链表头
	cur := root     // 充当游标

	for node1 != nil && node2 != nil {
		if node1.data < node2.data {
			cur.next = node1
			node1 = node1.next
		} else {
			cur.next = node2
			node2 = node2.next
		}
		cur = cur.next
	}

	if node1 != nil {
		cur.next = node1
	}
	if node2 != nil {
		cur.next = node2
	}

	return root.next
}

// 链表复制
func copyLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newNode := &Node{
		data: head.data + 1,
		next: nil,
	}

	curNode := head
	curNewNode := newNode

	for curNode.next != nil {
		curNode = curNode.next
		curNewNode.next = &Node{
			data: curNode.data,
		}
		curNewNode = curNewNode.next
	}
	return newNode
}

// 若链表有交集，则返回交点
func linkedListIntersection(node1, node2 *Node) *Node {
	if node1 == nil || node2 == nil {
		return nil
	}

	head2 := node2

	for node1 != nil {
		node2 = head2
		for node2 != nil {
			if node1.data == node2.data {
				return node1
			}
			node2 = node2.next
		}
		node1 = node1.next
	}

	return nil
}

// 判断链表是否有环(快慢指针)
func linkedListHasCycle(node *Node) bool {
	fast := node
	slow := node

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}

	return false
}

// 12 34 5 k = 2
// 21 43 5
func reverseKGroup(head *Node, k int) *Node {
	if head == nil || head.next == nil || k == 1 {
		return head
	}
	// 1、找到滑动窗口的头节点和尾节点
	tail := head
	count := k - 1
	for tail.next != nil {
		tail = tail.next
		count -= 1
		if count == 0 {
			reverseList(head).next = tail.next
			head = tail.next
			count = k
		}
	}
	return head
}

func reverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	var prev *Node
	cur := head

	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	return prev
}

func main() {
	head1 := &Node{0, nil}
	head2 := &Node{10, nil}

	for i := 1; i < 5; i++ {
		insert(head1, 0, 5-i)
	}

	for i := 1; i < 5; i++ {
		insert(head2, 0, 10-i)
	}

	//printElement(head1)
	//printElement(head2)

	fmt.Println(linkedListIntersection(head1, head2))
}
