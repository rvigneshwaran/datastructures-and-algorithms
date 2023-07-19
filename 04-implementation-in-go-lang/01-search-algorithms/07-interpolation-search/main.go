package main

import (
	"fmt"
	"math"
)

// Function to perform Interpolation Search
func interpolationSearch(arr []string, target string) int {
	low := 0                         // Initialize the low index of the search range
	high := len(arr) - 1             // Initialize the high index of the search range

	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		// Calculate the probable position of the target character using interpolation formula
		position := low + int(math.Floor(float64((target-arr[low])*(high-low))/(arr[high]-arr[low])))

		if arr[position] == target {
			return position
		}

		if arr[position] < target {
			low = position + 1
		} else {
			high = position - 1
		}
	}

	return -1
}

func main() {
	// Create a sorted slice of characters from Mahabharata
	characterList := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}

	// Perform Interpolation Search for a character
	character := "Draupadi"
	index := interpolationSearch(characterList, character)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", character, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
	}
}
