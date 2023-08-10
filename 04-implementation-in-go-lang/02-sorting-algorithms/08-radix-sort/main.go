package main

import (
	"fmt"
	"math/rand"
)

// Character represents a Harry Potter character with a name and magical power level
type Character struct {
	Name  string
	Power int
}

// getMaxPower function returns the maximum magical power level among characters
func getMaxPower(characters []Character) int {
	maxPower := 0
	for _, c := range characters {
		if c.Power > maxPower {
			maxPower = c.Power
		}
	}
	return maxPower
}

// countingSortByDigit function sorts characters based on a specific digit using Counting Sort
func countingSortByDigit(characters []Character, exp int) {
	n := len(characters)
	output := make([]Character, n)
	count := make([]int, 10) // 10 possible digits (0 to 9)

	// Count occurrences of each digit at the given exponent
	for i := 0; i < n; i++ {
		digit := (characters[i].Power / exp) % 10
		count[digit]++
	}

	// Update count array to store the actual positions in the output
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		digit := (characters[i].Power / exp) % 10
		output[count[digit]-1] = characters[i]
		count[digit]--
	}

	// Copy the sorted characters back to the original list
	copy(characters, output)
}

// radixSort function sorts the list of characters based on their magical power level using Radix Sort
func radixSort(characters []Character) {
	maxPower := getMaxPower(characters)

	// Perform counting sort for each digit position
	for exp := 1; maxPower/exp > 0; exp *= 10 {
		countingSortByDigit(characters, exp)
	}
}

// runTestSuite function runs a suite of test cases for Radix Sort
func runTestSuite() {
	fmt.Println("\nRunning Test Suite...")

	// Test Case 1: Sorting characters by power
	characters1 := []Character{
		{Name: "Luna Lovegood", Power: 250},
		{Name: "Draco Malfoy", Power: 180},
		{Name: "Ginny Weasley", Power: 310},
		{Name: "Neville Longbottom", Power: 220},
		{Name: "Dobby", Power: 130},
	}
	fmt.Println("\nTest Case 1:")
	fmt.Println("Unsorted List:")
	for _, c := range characters1 {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}
	radixSort(characters1)
	fmt.Println("Sorted List:")
	for _, c := range characters1 {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	// Test Case 2: Sorting characters with duplicate powers
	characters2 := []Character{
		{Name: "Sirius Black", Power: 220},
		{Name: "Remus Lupin", Power: 250},
		{Name: "Bellatrix Lestrange", Power: 220},
		{Name: "Severus Snape", Power: 250},
		{Name: "Lucius Malfoy", Power: 220},
	}
	fmt.Println("\nTest Case 2:")
	fmt.Println("Unsorted List:")
	for _, c := range characters2 {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}
	radixSort(characters2)
	fmt.Println("Sorted List:")
	for _, c := range characters2 {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	fmt.Println("\nTest Suite Completed!")
}

// generateLargeList function generates a large list of Harry Potter characters with random powers
func generateLargeList(size int) []Character {
	characters := make([]Character, size)
	for i := 0; i < size; i++ {
		characters[i] = Character{
			Name:  fmt.Sprintf("Character %d", i+1),
			Power: rand.Intn(1000), // Generating random powers between 0 and 999
		}
	}
	return characters
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

	// Perform Radix Sort
	radixSort(characters)

	fmt.Println("\nSorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	// Run the test suite
	runTestSuite()

	// Generate a large list of characters for testing
	largeList := generateLargeList(10000)

	// Perform Radix Sort on the large list
	radixSort(largeList)

	fmt.Println("\nSorted List:")
	for _, c := range largeList {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}
}
