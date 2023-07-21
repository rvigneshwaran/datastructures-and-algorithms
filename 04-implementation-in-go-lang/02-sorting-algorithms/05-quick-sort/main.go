package main

import (
	"fmt"
)

// Transformer represents a Transformer with a name and strength
type Transformer struct {
	Name     string
	Strength int
}

// quickSort function sorts the list of Transformers based on their strength in ascending order
func quickSort(transformers []Transformer) []Transformer {
	if len(transformers) <= 1 {
		return transformers
	}

	// Choose the pivot element (here, we choose the last element)
	pivotIndex := len(transformers) - 1
	pivot := transformers[pivotIndex]

	// Initialize left and right pointers
	left, right := 0, pivotIndex-1

	// Partition the list around the pivot element
	for left <= right {
		if transformers[left].Strength > pivot.Strength && transformers[right].Strength <= pivot.Strength {
			// Swap elements on the left and right sides of the pivot
			transformers[left], transformers[right] = transformers[right], transformers[left]
		}

		if transformers[left].Strength <= pivot.Strength {
			left++
		}

		if transformers[right].Strength > pivot.Strength {
			right--
		}
	}

	// Swap the pivot element to its correct position
	transformers[left], transformers[pivotIndex] = transformers[pivotIndex], transformers[left]

	// Recursively sort the two sublists
	quickSort(transformers[:left])
	quickSort(transformers[left+1:])

	return transformers
}

func main() {
	// List of Transformers (unsorted)
	transformers := []Transformer{
		{Name: "Optimus Prime", Strength: 100},
		{Name: "Megatron", Strength: 90},
		{Name: "Bumblebee", Strength: 80},
		{Name: "Starscream", Strength: 70},
		{Name: "Ironhide", Strength: 60},
		{Name: "Soundwave", Strength: 50},
	}

	fmt.Println("Unsorted List:")
	for _, t := range transformers {
		fmt.Printf("%s (Strength: %d)\n", t.Name, t.Strength)
	}

	// Perform Quick Sort
	transformers = quickSort(transformers)

	fmt.Println("\nSorted List:")
	for _, t := range transformers {
		fmt.Printf("%s (Strength: %d)\n", t.Name, t.Strength)
	}

	// Run the test suite
	runTestSuite()
}

// Function to run the test suite
func runTestSuite() {
	fmt.Println("\nRunning Test Suite:")

	// Test examples with different list sizes
	examples := []struct {
		Transformers []Transformer
	}{
		{Transformers: []Transformer{
			{Name: "Optimus Prime", Strength: 100},
			{Name: "Megatron", Strength: 90},
			{Name: "Bumblebee", Strength: 80},
			{Name: "Starscream", Strength: 70},
			{Name: "Ironhide", Strength: 60},
			{Name: "Soundwave", Strength: 50},
		}},
		{Transformers: generateLargeList(10000)}, // Large size list (10,000 elements)
	}

	for i, example := range examples {
		fmt.Printf("\nExample %d: Unsorted List\n", i+1)
		for _, t := range example.Transformers {
			fmt.Printf("%s (Strength: %d)\n", t.Name, t.Strength)
		}

		sortedTransformers := quickSort(example.Transformers)

		fmt.Printf("\nExample %d: Sorted List\n", i+1)
		for _, t := range sortedTransformers {
			fmt.Printf("%s (Strength: %d)\n", t.Name, t.Strength)
		}
	}
}

// Function to generate a large list of Transformers with random strengths
func generateLargeList(size int) []Transformer {
	transformers := make([]Transformer, size)
	for i := 0; i < size; i++ {
		transformers[i] = Transformer{Name: fmt.Sprintf("Transformer%d", i+1), Strength: i * 10}
	}
	return transformers
}
