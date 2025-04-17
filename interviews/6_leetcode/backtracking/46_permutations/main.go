package main

func bruteforce(nums []int, permutation []int, allPermutations *[][]int) {
    if len(nums) == 0 {
        // Добавляем текущую комбинацию в список всех перестановок
        permutationCopy := append([]int{}, permutation...)
        *allPermutations = append(*allPermutations, permutationCopy)
        return
    }

    n := len(nums)
    for i := 0; i < n; i++ {
        nextNum := nums[i]

        // Добавляем число в текущую перестановку и удаляем его из nums,
        // чтобы не использовать его повторно
        permutation = append(permutation, nextNum)
        nums[i], nums[n-1] = nums[n-1], nums[i]
        nums = nums[:n-1]

        // Рекурсивный вызов для оставшихся чисел
        bruteforce(nums, permutation, allPermutations)

        // Возвращаем состояние назад перед следующей итерацией
        nums = append(nums, nextNum)
        nums[i], nums[n-1] = nums[n-1], nums[i]
        permutation = permutation[:len(permutation)-1]
    }
}

func permutations(nums []int) [][]int {
    allPermutations := [][]int{}
    permutation := []int{}
    bruteforce(nums, permutation, &allPermutations)
    return allPermutations
}
