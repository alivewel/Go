package main

// totalCount = len(matrix) * len(matrix[0]) = 9
// m := r + l / 2 = 9 + 0 / 2 = 4
// i  = m / len(matrix) = 4 / 3 = 1
// j =  m % len(matrix) = 4 % 3 = 1

// [[1,2,3,4,5],
// [6,7,8,9,10],
// [11,12,13,14,15]]
// totalCount = len(matrix) * len(matrix[0]) = 15
// m := 8
// len(matrix) = 3; len(matrix[0] = 5.
// i  = m / len(matrix) = 8 / 3 = 2
// j =  m % len(matrix) = 8 % 3 = 2

// i  = m / len(matrix[0]) = 8 / 5 = 1
// j =  m % len(matrix[0]) - 1 = 8 % 5 - 1 = 3

// если m := 5


func search(matrix [][]int, target int) bool {

}