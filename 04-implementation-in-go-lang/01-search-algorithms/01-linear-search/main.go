package main

import "fmt"

// Function to perform linear search
func linearSearch(arr []string, key string) int {
	for i, value := range arr {
		if value == key {
			return i
		}
	}
	return -1
}

func main() {
	// Create a slice of characters from Mahabharata
	characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}

	// Perform linear search for a character
	character := "Draupadi"
	index := linearSearch(characters, character)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", character, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
	}
}