package main

import "fmt"

// Function to perform sentinel linear search
func searchCharacter(characters []string, target string) int {
	size := len(characters)
	lastElement := characters[size-1]
	characters[size-1] = target // Replace the last element with the target as sentinel

	index := 0
	for characters[index] != target {
		index++
	}

	characters[size-1] = lastElement // Restore the last element

	if index < size-1 || characters[size-1] == target {
		return index
	}

	return -1
}

func main() {
	// Create a slice of characters from Mahabharata
	characterList := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}

	// Perform sentinel linear search for a character
	targetCharacter := "Draupadi"
	index := searchCharacter(characterList, targetCharacter)

	// Check if the character was found or not
	if index != -1 {
		fmt.Printf("The character '%s' is found at index %d in the Mahabharata.\n", targetCharacter, index)
	} else {
		fmt.Printf("The character '%s' is not found in the Mahabharata.\n", targetCharacter)
	}
}