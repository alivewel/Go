package main

import (
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

    fmt.Println(node.Value) // Обработка текущего узла

    // Рекурсивный обход дочерних узлов
    for _, child := range node.Children {
        traverseTree(child)
    }
}

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