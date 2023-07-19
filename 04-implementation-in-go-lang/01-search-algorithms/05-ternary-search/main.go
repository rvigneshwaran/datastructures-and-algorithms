package main

import "fmt"

// Function to perform Ternary Search
func ternarySearch(arr []string, target string) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Divide the search range into three parts
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3
		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}
		if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}

	return -1
}

func main() {
	// Create a sorted slice of characters from Mahabharata
	characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}

	// Perform Ternary Search for a character
	character := "Draupadi"
	index := ternarySearch(characters, character)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", character, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
	}
}
