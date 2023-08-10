package main

import (
	"fmt"
	"math/rand"
)

// Character represents a Harry Potter character with a name and age
type Character struct {
	Name string
	Age  int
}

// countingSort function sorts the list of characters based on their age
func countingSort(characters []Character) {
	// Find the maximum age in the list
	maxAge := 0
	for _, c := range characters {
		if c.Age > maxAge {
			maxAge = c.Age
		}
	}

	// Create a count array and initialize it with zeros
	count := make([]int, maxAge+1)

	// Count the occurrences of each age
	for _, c := range characters {
		count[c.Age]++
	}

	// Update the count array to store the actual positions in the output
	for i := 1; i <= maxAge; i++ {
		count[i] += count[i-1]
	}

	// Create an output array and populate it using the count array
	output := make([]Character, len(characters))
	for i := len(characters) - 1; i >= 0; i-- {
		output[count[characters[i].Age]-1] = characters[i]
		count[characters[i].Age]--
	}

	// Copy the sorted characters back to the original list
	copy(characters, output)
}

func main() {
	// List of Harry Potter characters (unsorted)
	characters := []Character{
		{Name: "Harry Potter", Age: 18},
		{Name: "Hermione Granger", Age: 17},
		{Name: "Ron Weasley", Age: 18},
		{Name: "Severus Snape", Age: 38},
		{Name: "Voldemort", Age: 50},
	}

	fmt.Println("Unsorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}

	// Perform Counting Sort
	countingSort(characters)

	fmt.Println("\nSorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}

	// Generate a large list of characters for testing
	largeList := generateLargeList(10000)

	// Perform Counting Sort on the large list
	countingSort(largeList)

	fmt.Println("\nSorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}

	// Run the test suite
	runTestSuite()
}

// runTestSuite function runs a suite of test cases for the Counting Sort algorithm
func runTestSuite() {
	fmt.Println("\nRunning Test Suite...")

	// Test Case 1: Sorting characters by age
	characters1 := []Character{
		{Name: "Luna Lovegood", Age: 16},
		{Name: "Draco Malfoy", Age: 17},
		{Name: "Ginny Weasley", Age: 18},
		{Name: "Neville Longbottom", Age: 18},
		{Name: "Dobby", Age: 22},
	}
	fmt.Println("\nTest Case 1:")
	fmt.Println("Unsorted List:")
	for _, c := range characters1 {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}
	countingSort(characters1)
	fmt.Println("Sorted List:")
	for _, c := range characters1 {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}

	// Test Case 2: Sorting characters with duplicate ages
	characters2 := []Character{
		{Name: "Sirius Black", Age: 40},
		{Name: "Remus Lupin", Age: 38},
		{Name: "Bellatrix Lestrange", Age: 40},
		{Name: "Severus Snape", Age: 38},
		{Name: "Lucius Malfoy", Age: 40},
	}
	fmt.Println("\nTest Case 2:")
	fmt.Println("Unsorted List:")
	for _, c := range characters2 {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}
	countingSort(characters2)
	fmt.Println("Sorted List:")
	for _, c := range characters2 {
		fmt.Printf("%s (Age: %d)\n", c.Name, c.Age)
	}

	fmt.Println("\nTest Suite Completed!")
}

// generateLargeList function generates a large list of Harry Potter characters with random ages
func generateLargeList(size int) []Character {
	characters := make([]Character, size)
	for i := 0; i < size; i++ {
		characters[i] = Character{
			Name: fmt.Sprintf("Character %d", i+1),
			Age:  rand.Intn(101), // Generating random ages between 0 and 100
		}
	}
	return characters
}
