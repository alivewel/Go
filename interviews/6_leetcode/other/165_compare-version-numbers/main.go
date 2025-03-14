package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Atoi: Это сокращение от "ASCII to Integer".
// Itoa: Это сокращение от "Integer to ASCII".

func compareVersion(version1 string, version2 string) int {
	arrVer1 := strings.Split(version1, ".")
	arrVer2 := strings.Split(version2, ".")
	for i := range min(len(arrVer1), len(arrVer2)) {
		num1, _ := strconv.Atoi(arrVer1[i])
		num2, _ := strconv.Atoi(arrVer2[i])
		// fmt.Println(num1, num2)
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	for i := max(len(arrVer1), len(arrVer2))-1; i >= min(len(arrVer1), len(arrVer2)); i-- {
		var num int
		firstMore := true
        if len(arrVer1) > len(arrVer2) {
            num, _ = strconv.Atoi(arrVer1[i])

        } else {
            num, _ = strconv.Atoi(arrVer2[i])
			firstMore = false
        }
        // Если остались элементы, они должны быть равны нулю
        if num > 0 {
			if firstMore {
				return 1
			} else {
				return -1
			}
            
        }
	}
	return 0
}

func main() {
	// version1 := "1.01"
	// version2 := "1.001"
	// version1 := "1.0"
	// version2 := "1.0.0.0"
	version1 := "1"
	version2 := "1.1"
	fmt.Println(compareVersion(version1, version2))
}


// for i := range version1 {

// }
// p1, p2 := 0, 0
// if p1 
// for p1 + 1 < len() && version1[p1 + 1] != '.' {
	
// }