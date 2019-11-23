package binary_search

import (
	"fmt"
	"math"
)

func BinarySearch(target int, searchArray []int) int {
	searchArrayLength := len(searchArray)

	fmt.Printf("Searching %v for %d...\n", searchArray, target)

	// Check if we even need to search
	if searchArrayLength == 1 {
		if searchArray[0] == target {
			return 0
		}
		return -1
	}

	searchIndex := int(math.Ceil(float64(searchArrayLength) / float64(2)) - 1)
	searchValue := searchArray[searchIndex]

	if searchValue == target {
		return searchIndex
	} else if searchValue > target {
		return BinarySearch(target, searchArray[:searchIndex])
	} else if searchValue < target {
		subSearchIndex := BinarySearch(target, searchArray[searchIndex+ 1:])
		return subSearchIndex + searchIndex + 1
	}
	return -1
}
