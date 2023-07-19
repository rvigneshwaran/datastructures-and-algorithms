package main

import (
	"fmt"
)

// Transformer structure represents a Transformer
type Transformer struct {
	Name     string
	Strength int
}

// Function to perform Enhanced Binary Search on Transformers
func enhancedBinarySearch(transformers []Transformer, targetStrength int) int {
	low := 0
	high := len(transformers) - 1

	for low <= high {
		// Calculate mid using an enhanced formula to avoid integer overflow for large lists
		mid := low + (high-low)/2

		if transformers[mid].Strength == targetStrength {
			return mid // Target found at index mid
		} else if transformers[mid].Strength < targetStrength {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	// List of Transformers sorted by strength (for demonstration purposes)
	transformers := []Transformer{
		{Name: "Optimus Prime", Strength: 100},
		{Name: "Megatron", Strength: 90},
		{Name: "Bumblebee", Strength: 80},
		{Name: "Starscream", Strength: 70},
		{Name: "Ironhide", Strength: 60},
		{Name: "Soundwave", Strength: 50},
	}

	// Perform Enhanced Binary Search for a target strength
	targetStrength := 80
	index := enhancedBinarySearch(transformers, targetStrength)

	// Check if the target strength was found or not
	if index != -1 {
		fmt.Printf("Transformer with strength %d found at index %d.\n", targetStrength, index)
	} else {
		fmt.Printf("Transformer with strength %d not found.\n", targetStrength)
	}

	// Run the test suite
	runTestSuite()
}

// Function to run the test suite
func runTestSuite() {
	fmt.Println("Running Test Suite:")

	// Test examples with different list sizes
	examples := []struct {
		Transformers []Transformer
		Target       int
	}{
		{Transformers: []Transformer{{Name: "Optimus Prime", Strength: 100}}, Target: 100}, // Single element list
		{Transformers: []Transformer{
			{Name: "Optimus Prime", Strength: 100},
			{Name: "Megatron", Strength: 90},
			{Name: "Bumblebee", Strength: 80},
			{Name: "Starscream", Strength: 70},
			{Name: "Ironhide", Strength: 60},
			{Name: "Soundwave", Strength: 50},
		}, Target: 80}, // Medium size list
		{Transformers: generateLargeList(1000000), Target: 900000}, // Large size list (1 million elements)
	}

	for i, example := range examples {
		index := enhancedBinarySearch(example.Transformers, example.Target)

		if index != -1 {
			fmt.Printf("Example %d: Transformer with strength %d found at index %d.\n", i+1, example.Target, index)
		} else {
			fmt.Printf("Example %d: Transformer with strength %d not found.\n", i+1, example.Target)
		}
	}
}

// Function to generate a large list of Transformers with random strengths
func generateLargeList(size int) []Transformer {
	transformers := make([]Transformer, size)
	for i := 0; i < size; i++ {
		transformers[i] = Transformer{Name: fmt.Sprintf("Transformer%d", i+1), Strength: i * 1000}
	}
	return transformers
}
