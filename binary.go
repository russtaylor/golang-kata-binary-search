package binary_search

import (
	"fmt"
	"math"
)

func BinarySearch(target int, searchArray []int) int {
	searchArrayLength := len(searchArray)

	// Check if we even need to search
	if searchArrayLength == 1 {
		if searchArray[0] == target {
			return 0
		}
		return -1
	}

	searchTarget := int(math.Ceil(float64(searchArrayLength) / float64(2)) - 1)
	fmt.Println(searchTarget)
	if searchArray[searchTarget] == target {
		return searchTarget
	}
	return -1
}
