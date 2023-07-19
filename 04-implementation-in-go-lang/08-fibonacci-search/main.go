package main

import (
	"fmt"
	"time"
)

// Transformer structure represents a Transformer
type Transformer struct {
	Name     string
	Strength int
}

// Function to perform Fibonacci Search on Transformers
func fibonacciSearch(transformers []Transformer, targetStrength int) int {
	n := len(transformers)
	if n == 0 {
		return -1 // Empty list, target not found
	}

	// Initialize Fibonacci numbers
	fibPrev := 0
	fibCurr := 1
	fibNext := fibPrev + fibCurr

	// Find the smallest Fibonacci number greater than or equal to n
	for fibNext < n {
		fibPrev = fibCurr
		fibCurr = fibNext
		fibNext = fibPrev + fibCurr
	}

	// Initialize the search range
	offset := -1

	for fibNext > 1 {
		// Check if fibPrev is a valid index
		i := min(offset+fibPrev, n-1)

		if transformers[i].Strength == targetStrength {
			return i // Target found at index i
		} else if transformers[i].Strength < targetStrength {
			// Search in the right half
			fibNext = fibCurr
			fibCurr = fibPrev
			fibPrev = fibNext - fibCurr
			offset = i
		} else {
			// Search in the left half
			fibNext = fibPrev
			fibCurr = fibCurr - fibPrev
			fibPrev = fibNext - fibCurr
		}
	}

	// Check if the last element in the list is the target
	if fibCurr == 1 && offset+1 < n && transformers[offset+1].Strength == targetStrength {
		return offset + 1
	}

	return -1 // Target not found
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	// Perform Fibonacci Search for a target strength
	targetStrength := 80
	index := fibonacciSearch(transformers, targetStrength)

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
		startTime := time.Now()
		index := fibonacciSearch(example.Transformers, example.Target)
		elapsedTime := time.Since(startTime)

		if index != -1 {
			fmt.Printf("Example %d: Transformer with strength %d found at index %d. Time taken: %v\n", i+1, example.Target, index, elapsedTime)
		} else {
			fmt.Printf("Example %d: Transformer with strength %d not found. Time taken: %v\n", i+1, example.Target, elapsedTime)
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
