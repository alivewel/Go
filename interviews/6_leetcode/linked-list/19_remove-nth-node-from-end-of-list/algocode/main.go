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
	// Создание первого списка [1, 2, 3, 4, 5]
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}
	// list := &ListNode{Val:1, Next: &ListNode{2, nil}}


	// Создание второго списка [1, 2, 3, 5]
	res := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{5, nil}}}}

	fmt.Println(removeNthFromEnd(list, 5)) // Список: 8->9->9->9->0->0->0->1 Число: 10009998
	fmt.Println(res)
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy
    for i := 0; i < n + 1; i++ { // уточнить n + 1
        fast = fast.Next
    }
    for fast != nil { // условие уточнить
       fast = fast.Next 
       slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

// Решать задачу на яндекс код, а запускать на литкоде.
// Потому что литкод обесценивает цену запуска.
// Очень легко написать и запустить. На интервью такого не будет.

// Первый запуск бесплатно. Второй -10к. Третий -100к. От суммы оффера.
