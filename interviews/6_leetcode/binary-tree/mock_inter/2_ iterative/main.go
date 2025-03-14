package main

import (
    "container/list"
    "fmt"
)

// Определяем структуру узла дерева
type TreeNode struct {
    Value    string
    Children []*TreeNode
}

// Функция для обхода дерева
func traverseTree(node *TreeNode) {
    if node == nil {
        return // Проверка на наличие узла
    }

    res := make(string[], 0)

    queue := NewQueue()
    queue.Enqueue(node)

    // Рекурсивный обход дочерних узлов
    for !queue.IsEmpty() {
        node := queue.Dequeue()
        res = append(res, node.Value)
        for _, child := range root.children {
			queue.Enqueue(child)
		}
    }
    fmt.Println(res)
}

// func main() {
//     
    
//     queue.Enqueue("a")
//     queue.Enqueue("b")
//     queue.Enqueue("c")

//     fmt.Println("Size of queue:", queue.Size()) // Output: Size of queue: 3

//     for !queue.IsEmpty() {
//         fmt.Println(queue.Dequeue()) // Output: a b c
//     }
// }

func main() {
    // Создаем структуру дерева
    tree := &TreeNode{
        Value: "a",
        Children: []*TreeNode{
            {
                Value: "b",
                Children: []*TreeNode{
                    {
                        Value: "d",
                        Children: []*TreeNode{
                            {
                                Value: "h",
                                Children: []*TreeNode{
                                    {
                                        Value: "i",
                                        Children: []*TreeNode{},
                                    },
                                },
                            },
                        },
                    },
                    {
                        Value: "e",
                        Children: []*TreeNode{},
                    },
                },
            },
            {
                Value: "c",
                Children: []*TreeNode{
                    {
                        Value: "f",
                        Children: []*TreeNode{},
                    },
                    {
                        Value: "g",
                        Children: []*TreeNode{},
                    },
                },
            },
        },
    }

    // Запускаем обход дерева
    traverseTree(tree)
}




// Queue представляет собой структуру очереди
type Queue struct {
    items *list.List
}

// NewQueue создает новую очередь
func NewQueue() *Queue {
    return &Queue{
        items: list.New(),
    }
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(item interface{}) {
    q.items.PushBack(item)
}

// Dequeue удаляет и возвращает элемент из начала очереди
func (q *Queue) Dequeue() interface{} {
    if q.items.Len() == 0 {
        return nil // Возвращаем nil, если очередь пуста
    }
    front := q.items.Front()
    q.items.Remove(front)
    return front.Value
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
    return q.items.Len() == 0
}

// Size возвращает количество элементов в очереди
func (q *Queue) Size() int {
    return q.items.Len()
}
