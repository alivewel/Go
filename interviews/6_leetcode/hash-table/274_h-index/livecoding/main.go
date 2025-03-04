package main

// citations=[10,3,5,0,3]
//countCit=[0, 0,0,2,0,2]


//countCit=[0, 1,0,1,0,2]
// count
// arr[len(citations+1)] = 
// time O(2n) => O(n)
// mem O(n)

func findHirschIndex(citations []int) int {
    countCit := make([]int, len(citations)+1)
    for _, val := range citations {
        if val > len(citations) {
            countCit[len(citations)]++
        } else {
            countCit[val]++
        }
    }
    hIndex := 0
    //countCit=[0, 0,0,2,0,2]
    for i := len(countCit)-1; i >= 0; i-- {
        hIndex += countCit[i]
        // 3    4
        if i <= hIndex {
            return i
        }
    }
    return 0
}

func main() {
	nums := []int{10, 1, 8, 0, 3}
	fmt.Println(findHirschIndex(nums))
}
