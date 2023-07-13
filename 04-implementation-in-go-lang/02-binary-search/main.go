package main

import "fmt"

// Function to perform binary search
func binarySearch(arr []string, key string) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	// Create a sorted slice of characters from Mahabharata
	characters := []string{"Arjuna", "Bhima", "Draupadi", "Duryodhana", "Krishna", "Nakula", "Sahadeva", "Yudhishthira"}

	// Perform binary search for a character
	character := "Krishna"
	index := binarySearch(characters, character)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", character, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
	}
}
