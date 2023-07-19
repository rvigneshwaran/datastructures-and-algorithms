package main

import "fmt"

// Function to perform Meta Binary Search
func metaBinarySearch(arr [][]string, target string) (int, int) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		midElement := arr[mid][0]

		if midElement == target {
			return mid, 0
		} else if midElement < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high, low
}

// Function to perform Binary Search within a range
func binarySearch(arr []string, target string, low int, high int) int {
	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	// Create a slice of sorted slices, each representing characters from a specific section of Mahabharata
	sections := [][]string{
		{"Adi Parva", "Amba", "Bhishma", "Chitrangada"},
		{"Karna Parva", "Karna", "Kunti"},
		{"Virata Parva", "Draupadi", "Arjuna"},
		{"Udyoga Parva", "Yudhishthira", "Duryodhana"},
	}

	// Perform Meta Binary Search for a character
	character := "Arjuna"
	sectionIndex, offset := metaBinarySearch(sections, character)

	// Check if the character was found or not
	if sectionIndex != -1 {
		characters := sections[sectionIndex]
		characterIndex := binarySearch(characters, character, 1, len(characters)-1)

		if characterIndex != -1 {
			fmt.Printf("The character '%s' is found in the section '%s' at index %d in the Mahabharata.\n", character, sections[sectionIndex][0], characterIndex)
		} else {
			fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
		}
	} else {
		if offset == 0 {
			fmt.Printf("The character '%s' is found in the section '%s' in the Mahabharata.\n", character, sections[0][0])
		} else {
			fmt.Printf("The character '%s' is not found in the Mahabharata.\n", character)
		}
	}
}
