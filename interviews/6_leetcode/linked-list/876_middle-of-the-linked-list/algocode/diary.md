1) С таким решением выходит паника.
Нужно 2 условия в цикле: fast != nil && fast.Next != nil 

```
func middleNodeValue(head *ListNode) int {
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        fast = slow
        if fast.Next != nil {
            fast = fast.Next
        }
    }
    return slow.Val
}
```

2) 
Имеется:
fast = slow.Next
Должно быть:
fast = fast.Next.Next

```
func middleNodeValue(head *ListNode) int {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
		fast = fast.Next.Next
    }
    return slow.Val
}
```