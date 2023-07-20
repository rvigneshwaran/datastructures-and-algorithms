package main

import "fmt"

// Transformer represents a Transformer with a name and strength
type Transformer struct {
	Name     string
	Strength int
}

// insertionSort function sorts the list of Transformers based on their strength in ascending order
func insertionSort(transformers []Transformer) {
	for i := 1; i < len(transformers); i++ {
		key := transformers[i]
		j := i - 1

		// Move elements greater than the key to one position ahead of their current position
		for j >= 0 && transformers[j].Strength > key.Strength {
			transformers[j+1] = transformers[j]
			j--
		}

		// Insert the key in its correct position
		transformers[j+1] = key
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

	// Perform Insertion Sort
	insertionSort(transformers)

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

		insertionSort(example.Transformers)

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
