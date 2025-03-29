package main

// [[index, value], [index, value]]
// [5,6,9,7,5,-1,8,11,2]
// [1,1,5,3,2,1,1,0,0]

func predictWarming(temperatures []int) []int {
    stack := make([][]int, 0)
    // stack = append(stack, []int{0, temperatures[0]})
    result := make([]int, len(temperatures))
    for i := 0; i < len(temperatures); i++ { // i := 1
        for len(stack) > 0 && stack[len(stack)-1][1] < temperatures[i] {
            result[stack[len(stack)-1][0]] = i-stack[len(stack)-1][0]
            stack = stack[:len(stack)-1] // stack = stack[:0]
        }
        stack = append(stack, []int{i, temperatures[i]})
    }
    return result
}
