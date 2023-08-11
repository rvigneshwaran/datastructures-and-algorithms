package main

import (
	"fmt"
	"sort"
)

// Character represents a Harry Potter character with a name and magical power level
type Character struct {
	Name  string
	Power int
}

// ByPower implements the sort.Interface for []Character based on Power
type ByPower []Character

func (a ByPower) Len() int           { return len(a) }
func (a ByPower) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPower) Less(i, j int) bool { return a[i].Power < a[j].Power }

// displayCharacters function displays the list of characters
func displayCharacters(characters []Character) {
	for _, c := range characters {
		fmt.Printf("%s (Power: %d)\n", c.Name, c.Power)
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
	displayCharacters(characters)

	// Perform Tim Sort using the standard Go sort package
	sort.Sort(ByPower(characters))

	fmt.Println("\nSorted List:")
	displayCharacters(characters)
}
