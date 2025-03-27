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
	// list1 := &ListNode{1,
	// 	&ListNode{2,
	// 		&ListNode{3,
	// 			&ListNode{4,
	// 				&ListNode{5, nil}}}}}

	// Создание списка [1, 2, 3, 4, 5, 6]
	// list2 := &ListNode{1,
	// 	&ListNode{2,
	// 		&ListNode{3,
	// 			&ListNode{4,
	// 				&ListNode{5,
	// 					&ListNode{6, nil}}}}}}

	// Создание списка [1, 2, 2, 1]
	list3 := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{1, nil}}}}

	// fmt.Println(isPalindrome(list1))
	// fmt.Println(isPalindrome(list2))
	fmt.Println(isPalindrome(list3))
}

// сначала найдем middleNode
// потом reverseList
// потом пройдем по обоим спискам, до тех пор один из них не равен null и сравним значения

func isPalindrome(head *ListNode) bool {
	middleHead := middleNode(head)
	fmt.Println(middleHead)
	newHead := reverseList(middleHead)
	fmt.Println(newHead)
	for newHead != nil {
		if newHead.Val != head.Val {
			return false
		}
		newHead = newHead.Next
		head = head.Next
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 3 указателя
// текущий
// предыдущий
// следующий

// меняем указатели

// внутри цикла создаем временный от текущего
// текущий делаем следующим
// предыдущий приравнивем к текущему
// у текущего следующий ставим на предыдущий

func reverseList(head *ListNode) *ListNode {
	// prev := &ListNode{} // создается лишний нулевой узел в конце
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr
		curr = curr.Next
		temp.Next = prev
		prev = temp
	}
	return prev
}
