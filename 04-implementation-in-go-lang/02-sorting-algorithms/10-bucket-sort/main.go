package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Character represents a Harry Potter character with a name and magical power level
type Character struct {
	Name  string
	Power int
}

// bucketSort function sorts the list of characters based on their magical power level using Bucket Sort
func bucketSort(characters []Character) {
	n := len(characters)
	if n <= 1 {
		return // Nothing to sort
	}

	// Determine the maximum power level
	maxPower := 0
	for _, c := range characters {
		if c.Power > maxPower {
			maxPower = c.Power
		}
	}

	// Create buckets and distribute characters into them
	numBuckets := n
	buckets := make([][]Character, numBuckets)
	for _, c := range characters {
		index := (c.Power * (numBuckets - 1)) / maxPower
		buckets[index] = append(buckets[index], c)
	}

	// Sort characters within each bucket
	for _, bucket := range buckets {
		sort.SliceStable(bucket, func(i, j int) bool {
			return bucket[i].Power < bucket[j].Power
		})
	}

	// Concatenate sorted buckets to get the final sorted list
	index := 0
	for _, bucket := range buckets {
		copy(characters[index:], bucket)
		index += len(bucket)
	}
}

// runTestSuite function runs a suite of test cases for Bucket Sort
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
	bucketSort(characters1)
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
	bucketSort(characters2)
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

	// Perform Bucket Sort
	bucketSort(characters)

	fmt.Println("\nSorted List:")
	for _, c := range characters {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	// Run the test suite
	runTestSuite()

	// Generate a large list of characters for testing
	largeList := generateLargeList(10000)

	// Perform Bucket Sort on the large list
	bucketSort(largeList)

	// Display and validate results on the large list
	fmt.Println("\nLarge List (After Bucket Sort):")
	for _, c := range largeList {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
	}

	// Validate the sorted order on the large list
	fmt.Println("\nValidating Sorted Order...")
	isSorted := true
	for i := 1; i < len(largeList); i++ {
		if largeList[i].Power < largeList[i-1].Power {
			isSorted = false
			break
		}
	}
	if isSorted {
		fmt.Println("Sorted order is validated.")
	} else {
		fmt.Println("Sorted order is NOT validated.")
	}
}
