package main

import (
	"fmt"
)

// Character represents a Harry Potter character with a name and magical power level
type Character struct {
	Name  string
	Power int
}

// shellSort function sorts the list of characters based on their magical power level using Shell Sort
func shellSort(characters []Character) {
	n := len(characters)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := characters[i]
			j := i

			for j >= gap && characters[j-gap].Power > temp.Power {
				characters[j] = characters[j-gap]
				j -= gap
			}
			characters[j] = temp
		}
		gap /= 2
	}
}

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

	// Perform Shell Sort
	shellSort(characters)

	fmt.Println("\nSorted List:")
	displayCharacters(characters)
}