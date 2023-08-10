package main

import (
	"fmt"
)

// Character represents a Harry Potter character with a name and magical power level
type Character struct {
	Name  string
	Power int
}

// heapSort function sorts the list of characters based on their magical power level in descending order
func heapSort(characters []Character) {
	buildMaxHeap(characters)
	for i := len(characters) - 1; i >= 1; i-- {
		// Move the maximum element (root) to the end of the list
		characters[0], characters[i] = characters[i], characters[0]
		maxHeapify(characters[:i], 0)
	}
}

// buildMaxHeap function builds a max heap from the list of characters
func buildMaxHeap(characters []Character) {
	for i := len(characters)/2 - 1; i >= 0; i-- {
		maxHeapify(characters, i)
	}
}

// maxHeapify function re-heapifies the heap rooted at index i
func maxHeapify(characters []Character, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	// Find the largest element among the root, left child, and right child
	if left < len(characters) && characters[left].Power > characters[largest].Power {
		largest = left
	}
	if right < len(characters) && characters[right].Power > characters[largest].Power {
		largest = right
	}

	// If the largest element is not the root, swap the elements and continue maxHeapify
	if largest != i {
		characters[i], characters[largest] = characters[largest], characters[i]
		maxHeapify(characters, largest)
	}
}

func main() {
	// List of Harry Potter characters (unsorted)
	characters := []Character{
		{Name: "Harry Potter", Power: 95},
		{Name: "Hermione Granger", Power: 90},
		{Name: "Ron Weasley", Power: 85},
		{Name: "Severus Snape", Power: 80},
		{Name: "Voldemort", Power: 100},
	}

	fmt.Println("Unsorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	// Perform Heap Sort
	heapSort(characters)

	fmt.Println("\nSorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

}
