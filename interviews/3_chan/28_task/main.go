package main

import (
	"fmt"
	"strconv"
)

// 2. Add Two Numbers

type Item struct {
	val  int
	next *Item
}

func (it *Item) String() (result string) {
	for cur := it; cur != nil; cur = cur.next {
		if cur == cur.next {
			return "CYCLE DETECTED"
		}
		// result = strconv.Itoa(cur.val) + " " + result
		result += strconv.Itoa(cur.val) + " "
	}
	return result
}

func main() {

	// Список: 2->4->3 Число:342
	first := &Item{2, &Item{4, &Item{3, nil}}}
	// Список: 5->6->4 Число:465
	second := &Item{5, &Item{6, &Item{4, nil}}}

	// zero := &Item{0, nil}
	// // Список: 9->9->9->9->9->9->9 Число: 9999999
	// third := &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, &Item{9, nil}}}}}}}
	// // Список: 9->9->9->9 Число: 9999
	// fourth := &Item{9, &Item{9, &Item{9, &Item{9, nil}}}}

	fmt.Println(addTwoNumbers(first, second)) // Список: 7->0->8  Число: 342+465=807
	// fmt.Println(addTwoNumbers(zero, zero))    // 0
	// fmt.Println(addTwoNumbers(third, fourth)) // Список: 8->9->9->9->0->0->0->1 Число: 10009998

}

func addTwoNumbers(a, b *Item) *Item {
	res := &Item{}
	head := res
	temp := 0
	over := false
	for a != nil && b != nil {
		temp = a.val + b.val
		if over {
			temp += 1
			over = false
		}
		if temp > 9 {
			over = true
			temp = temp % 10
		}
		res.val = temp
		a = a.next
		b = b.next

		if a != nil || b != nil {
			res.next = &Item{}
			res = res.next
		}
	}

	if a != nil {
		res.next = b
	} else if b != nil {
		res.next = a
	}

	return head
}
