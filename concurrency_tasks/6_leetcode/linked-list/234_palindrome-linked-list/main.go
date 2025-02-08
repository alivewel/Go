package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (it *ListNode) String() (result string) {
	for cur := it; cur != nil; cur = cur.Next {
		if cur == cur.Next {
			return "CYCLE DETECTED"
		}
		result += strconv.Itoa(cur.Val) + " "
	}
	return result
}

func main() {
	// Создание списка [1, 2, 3, 4, 5]
	list1 := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}

	// Создание списка [1, 2, 3, 4, 5, 6]
	list2 := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						&ListNode{6, nil}}}}}}

	// Создание списка [1, 2, 2, 1]
	list3 := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{1, nil}}}}

	fmt.Println(isPalindrome(list1))
	fmt.Println(isPalindrome(list2))
	fmt.Println(isPalindrome(list3))
}

func isPalindrome(head *ListNode) bool {
	firstHalfEndCode := middleNode(head)
	secondHalfBeginCode := reverseList(firstHalfEndCode)
	p1 := head
	p2 := secondHalfBeginCode
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}
