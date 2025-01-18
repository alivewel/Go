package main

import (
	"fmt"
	"math"
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
	list3 := &ListNode{1,
		&ListNode{2,
			&ListNode{11,
				&ListNode{21,
					&ListNode{25,
						&ListNode{31, nil}}}}}}

	fmt.Println(mergeKLists([]*ListNode{list1, list2, list3}))
}

func getVal(list *ListNode) int {
	if list == nil {
		return math.MaxInt
	}
	return list.Val
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	for list1 != nil || list2 != nil {
		if getVal(list1) > getVal(list2) {
			dummy.Next = list2
			list2 = list2.Next
		} else {
			dummy.Next = list1
			list1 = list1.Next
		}
		dummy = dummy.Next
	}
	return res.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	for {
		minIdx := getMinNodeIdx(lists)
		if minIdx == -1 {
			dummy.Next = nil
			break
		}
		dummy.Next = lists[minIdx]
		dummy = dummy.Next
		lists[minIdx] = lists[minIdx].Next
	}
	return res.Next
}

func getMinNodeIdx(lists []*ListNode) int {
	minVal, minIdx := math.MaxInt, math.MaxInt
	for i, curNode := range lists {
		if curNode == nil {
			continue
		}
		if curNode.Val < minVal {
			minVal = curNode.Val
			minIdx = i
		}
	}
	if minIdx == math.MaxInt {
		return -1
	}
	return minIdx
}

// Написать с нуля все функции
