package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
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
	head := &ListNode{}
	current := head
	// создать список
	// заполнить его значениями
	// и распечать
	for i := 0; i < 5; i++ {
		current.Val = i
		if i < 4 { // Не создаем следующий узел для последнего элемента
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	fmt.Println(head.Next)
}
