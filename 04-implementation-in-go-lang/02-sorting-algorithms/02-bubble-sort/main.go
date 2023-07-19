package main

import "fmt"

// Transformer represents a Transformer with a name and strength
type Transformer struct {
	Name     string
	Strength int
}

// bubbleSort function sorts the list of Transformers based on their strength in ascending order
func bubbleSort(transformers []Transformer) {
	n := len(transformers)
	swapped := true

	for swapped {
		swapped = false

		// Compare adjacent elements and swap if needed
		for i := 0; i < n-1; i++ {
			if transformers[i].Strength > transformers[i+1].Strength {
				transformers[i], transformers[i+1] = transformers[i+1], transformers[i]
				swapped = true
			}
		}

		// The last element in the iteration is the largest and in its correct position,
		// so we don't need to consider it in the next iteration.
		n--
	}
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

	// Perform Bubble Sort
	bubbleSort(transformers)

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

		bubbleSort(example.Transformers)

		fmt.Printf("\nExample %d: Sorted List\n", i+1)
		for _, t := range example.Transformers {
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
