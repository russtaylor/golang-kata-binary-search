package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, 0, 0)
}

func TestBinarySearchReturnsNegativeOneIfTargetNotFound(t *testing.T) {
	// Arrange
	inputArray := []int{3, 4, 5}
	searchTarget := 1
	expected := -1

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns0IfOnlyElementInArrayMatches(t *testing.T) {
	// Arrange
	inputArray := []int{3}
	searchTarget := 3
	expected := 0

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturnsMiddleIndexIfMiddleMatches(t *testing.T) {
	// Arrange
	inputArray := []int{1, 3, 5}
	searchTarget := 3
	expected := 1

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}
