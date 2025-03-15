/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

 1 -> 2 -> 3 -> 4
      7 -> 6 -> 9

 2  ->    0  ->   0  ->  3
 */

// 1. 2 числа переворачиваем через ReverseLindkedList
// 2. Создаем новый список и складываем 2 числа
// 3. Результат переворачиваем.
//  h <- n
//  1 -> 2 -> 3 <- 0
// сохранить указатель на next

func reverseLindkedList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil {
        nextNode := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextNode
    }
    return prev
}
// 222
//  79
// 301

func addTwoNumbers(num1, num2 *ListNode) *ListNode {
    num1 = reverseLindkedList(num1)
    num2 = reverseLindkedList(num2)
    var res *ListNode
    carry := 0
    for num1 != nil && num2 != nil {
        sum := num1.Val + num2.Val + carry
        carry = 0
        if sum > 9 {
            carry = sum / 10
        }  
        res.Next = &ListNode{Val: (sum % 10)} // 
        num1 = num1.Next
        num2 = num2.Next
    }
    for num1 != nil {
        res.Next = &ListNode{Val: num1.Val + carry} 
        num1 = num1.Next
        carry = 0
    }
    for num2 != nil {
        res.Next = &ListNode{Val: num2.Val + carry}
        num2 = num2.Next
        carry = 0
    }
    return res
}





