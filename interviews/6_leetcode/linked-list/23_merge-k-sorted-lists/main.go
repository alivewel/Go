package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}

// Функция для вывода связанного списка
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
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
	// Создаем три отсортированных связанных списка
	list1 := createList([]int{1, 4, 5}) // 1 -> 4 -> 5
	list2 := createList([]int{1, 3, 4}) // 1 -> 3 -> 4
	list3 := createList([]int{2, 6})    // 2 -> 6

	// Помещаем их в массив
	lists := []*ListNode{list1, list2, list3}

	// Теперь можно вызвать функцию mergeKLists
	mergedList := mergeKLists(lists)

	// Для вывода результата можно использовать следующую функцию
	printList(mergedList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for list1 != nil && list2 != nil {
		dummy := &ListNode{}
		if list2.Val > list1.Val {
			dummy = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			dummy = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		res.Next = dummy
		res = res.Next
	}
	if list1 != nil {
		res.Next = list1
	}
	if list2 != nil {
		res.Next = list2
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	fmt.Println("lists[:num]", lists[:num])
	fmt.Println("lists[num:]", lists[num:])
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

// у нас изначально имеется 3 односвязных списка
// list1: 1 -> 4 -> 5
// list2: 1 -> 3 -> 4
// list3: 2 -> 6

// при первом вызове функции mergeKLists
// мы получаем num := length / 2 равным 1
// соответственно получаем в левой части list1,
// а в правой list2 и list3
// вызываем mergeKLists рекрусивно снова
// так как в левой части length == 1 возвращаем сам list1 (return lists[0])
// а в правой части снова делим наш список на две части
// и получаем left == list2, right == list3
// вызываем mergeTwoLists() и сортируем 2 списка
// и возвращаем отсортированный список из list2 и list3
// далее рекурсивно возвращаемся вверх и возвращаем в right наш отсортированный список
// теперь вызываем mergeTwoLists() и передаем в него list1 и отсортированный список из list2 и list3
// получаем итоговый отсортированный список
