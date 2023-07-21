package main

import (
	"fmt"
)

// Transformer represents a Transformer with a name and strength
type Transformer struct {
	Name     string
	Strength int
}

// merge function merges two sorted halves into a single sorted list
func merge(left, right []Transformer) []Transformer {
	result := make([]Transformer, 0)
	l, r := 0, 0

	// Compare elements from both halves and add them to the result in sorted order
	for l < len(left) && r < len(right) {
		if left[l].Strength <= right[r].Strength {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Add any remaining elements from the left or right half to the result
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// mergeSort function sorts the list of Transformers based on their strength in ascending order
func mergeSort(transformers []Transformer) []Transformer {
	if len(transformers) <= 1 {
		return transformers
	}

	// Divide the list into two halves
	mid := len(transformers) / 2
	left := mergeSort(transformers[:mid])
	right := mergeSort(transformers[mid:])

	// Merge the two sorted halves
	return merge(left, right)
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

	// Perform Merge Sort
	transformers = mergeSort(transformers)

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

		sortedTransformers := mergeSort(example.Transformers)

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
