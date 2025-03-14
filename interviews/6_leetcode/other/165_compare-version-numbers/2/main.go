package main

import (
	"fmt"
	"strings"
	"strconv"
)

//
func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	n, m := len(arr1), len(arr2)
	// добавляем в меньший массив нули и выравниваем длины двух массив
	// одинаковые массивы проще сравнивать
	if n > m {
		for i := 0; i < n-m; i++ {
			arr2 = append(arr2, "0")
		}
	} else if n < m {
		for i := 0; i < m-n; i++ {
			arr1 = append(arr1, "0")
		}
	}

	for i := 0; i < len(arr1); i++ {
		one, _ := strconv.Atoi(arr1[i])
		two, _ := strconv.Atoi(arr2[i])
		if one > two {
			return 1
		} else if one < two {
			return -1
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
