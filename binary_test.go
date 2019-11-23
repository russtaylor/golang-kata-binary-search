package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, 0, 0)
}

func TestBinarySearchReturnsNegative1IfTargetNotFound(t *testing.T) {
	// Arrange
	inputArray := []int{3, 4, 5}
	searchTarget := 1
	expected := -1

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns0In1ElementArrayWithTarget(t *testing.T) {
	// Arrange
	inputArray := []int{3}
	searchTarget := 3
	expected := 0

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns1In3ElementArrayWhereMiddleIsTarget(t *testing.T) {
	// Arrange
	inputArray := []int{1, 3, 5}
	searchTarget := 3
	expected := 1

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns0In3ElementArrayWhereFirstElementIsTarget(t *testing.T) {
	// Arrange
	inputArray := []int{1, 3, 5}
	searchTarget := 1
	expected := 0

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns2In3ElementArrayWhereLastElementIsTarget(t *testing.T) {
	// Arrange
	inputArray := []int{1, 3, 5}
	searchTarget := 5
	expected := 2

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestBinarySearchReturns2In8ElementArrayWhereThirdElementIsTarget(t *testing.T) {
	// Arrange
	inputArray := []int{1, 3, 5, 11, 32, 42, 64, 210}
	searchTarget := 5
	expected := 2

	// Act
	actual := BinarySearch(searchTarget, inputArray)

	// Assert
	assert.Equal(t, expected, actual)
}

// TODO Check for unsorted array
