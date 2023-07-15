package main

import (
	"fmt"
	"math"
)

// Function to perform Jump Search
func jumpSearch(arr []string, target string) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n))) // Determine the block size

	prev := 0 // Previous index

	// Finding the block where the target element might be present
	for arr[int(math.Min(float64(step), float64(n)))-1] < target { // Jumping ahead until the value at index `step` or end of array is greater than target
		prev = step   // Store previous index
		step += int(math.Sqrt(float64(n))) // Increment the step

		if prev >= n { // If previous index exceeds array size, target not found
			return -1
		}
	}

	// Performing linear search within the identified block
	for arr[prev] < target { // Linear search within the identified block
		prev++ // Move to the next element in the block

		if prev == int(math.Min(float64(step), float64(n))) { // If end of block is reached, target not found
			return -1
		}
	}

	// Checking if the target element is found
	if arr[prev] == target { // If target element found, return its index
		return prev
	}

	return -1 // Target element not found
}

func main() {
	// Create a sorted slice of characters from Mahabharata
	characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}

	// Perform Jump Search for a character
	character := "Draupadi"
	index := jumpSearch(characters, character)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", character, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
	}
}
