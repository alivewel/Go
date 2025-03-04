
package main

import (
	"fmt"
	"strconv"
)

// [0,0,0,1,1,1,2,2,3,3,4]
// time O(n)
// mem O(1)
func removeDuplicates(nums []int) int {
    slow, fast := 0, 0
    k := 0
    for fast < len(nums) {
        for fast+1 < len(nums) && nums[fast] == nums[fast+1] {
            fast++
        }
        // if slow != fast {
            nums[k] = nums[slow]
        // }
        slow = fast + 1
        fast = fast + 1
        k++
    }
    return k
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4} //
	fmt.Println(removeDuplicates(nums))
}
