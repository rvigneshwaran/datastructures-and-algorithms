# Meta Binary Search Algorithm in Go

This repository contains an implementation of the Meta Binary Search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.


## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Explanation](#explanation)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The Meta Binary Search algorithm is a combination of meta search and binary search. It is used to search for a target element within specific sections of a larger dataset. The algorithm follows these steps:

1. Perform a binary search (meta search) on the sections of the dataset to find the section that may contain the target element.
2. If the section is found, perform a binary search within that section to find the exact index of the target element.
3. Return the section index and the element index if found, or indicate that the target element was not found.

The time complexity of the Meta Binary Search algorithm depends on the number of sections (n) and the number of elements in the identified section (m). The meta search has a time complexity of O(log n), and the binary search has a time complexity of O(log m).

## Usage

The main program file (`main.go`) contains an example of how to use the Meta Binary Search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Explanation

In the implemented code, the Meta Binary Search algorithm is applied to search for a target character in the Mahabharata character list.

1. The `metaBinarySearch` function performs the meta search on a slice of sorted slices (`arr`) to find the section that may contain the target character. It returns the index of the section (if found) and an offset indicating whether the target character is found in the section.
2. The `binarySearch` function performs the binary search within a given range of a slice to find the exact index of the target character.
3. In the `main` function, a slice of sorted slices (`sections`) is created, where each sub-slice represents characters from a specific section of the Mahabharata.
4. The `metaBinarySearch` function is called to find the section that may contain the character.
5. If the section is found, the `binarySearch` function is called within that section to find the exact index of the character.
6. Finally, the program prints whether the character was found or not, along with the section name and the character's index (if found).

## Examples
The program includes multiple examples to demonstrate the Meta Binary search algorithm and its time complexity.

Example 1:
```go
sections := [][]string{
	{"Adi Parva", "Amba", "Bhishma", "Chitrangada"},
	{"Karna Parva", "Karna", "Kunti"},
	{"Virata Parva", "Draupadi", "Arjuna"},
	{"Udyoga Parva", "Yudhishthira", "Duryodhana"},
}
character := "Arjuna"
```
Output : 
```shell
The character 'Arjuna' is found in the section 'Virata Parva' at index 2 in the Mahabharata.
```

In this example, the target character "Arjuna" is found in the section "Virata Parva" at index 2. The Meta Binary Search algorithm starts with a meta search on the sections to identify the potential section containing the target character. It performs a binary search on the sections and determines that the section "Virata Parva" might contain the target character. The subsequent binary search within the identified section finds the target character at index 2.

The time complexity of the Meta Binary Search algorithm is determined by the number of sections (n) and the number of elements in the identified section (m). In this example, the meta search involves performing a binary search on four sections, resulting in a time complexity of O(log n). The subsequent binary search within the identified section has a time complexity of O(log m). Therefore, the overall time complexity for this example is O(log n + log m).

Example 2:
```go
sections := [][]string{
	{"Adi Parva", "Amba", "Bhishma", "Chitrangada"},
	{"Karna Parva", "Karna", "Kunti"},
	{"Virata Parva", "Draupadi", "Arjuna"},
	{"Udyoga Parva", "Yudhishthira", "Duryodhana"},
}
character := "Karna"
```
Output : 
```shell
The character 'Karna' is found in the section 'Karna Parva' at index 1 in the Mahabharata.
```

In this example, the target character "Karna" is found in the section "Karna Parva" at index 1. The Meta Binary Search algorithm performs a meta search on the sections and identifies the section "Karna Parva" as the potential section containing the target character. The subsequent binary search within the identified section finds the target character at index 1.

Similar to the previous example, the time complexity is determined by the number of sections (n) and the number of elements in the identified section (m). In this case, the meta search still involves performing a binary search on four sections, resulting in a time complexity of O(log n). The subsequent binary search within the identified section has a time complexity of O(log m). Therefore, the overall time complexity for this example is O(log n + log m).

Example 3:
```go
sections := [][]string{
	{"Adi Parva", "Amba", "Bhishma", "Chitrangada"},
	{"Karna Parva", "Karna", "Kunti"},
	{"Virata Parva", "Draupadi", "Arjuna"},
	{"Udyoga Parva", "Yudhishthira", "Duryodhana"},
}
character := "Duryodhana"
```
Output : 
```shell
The character 'Duryodhana' is found in the section 'Udyoga Parva' at index 2 in the Mahabharata.
```
In this example, the target character "Duryodhana" is found in the section "Udyoga Parva" at index 2. The Meta Binary Search algorithm performs a meta search on the sections and identifies the section "Udyoga Parva" as the potential section containing the target character. The subsequent binary search within the identified section finds the target character at index 2.

Similar to the previous examples, the time complexity is determined by the number of sections (n) and the number of elements in the identified section (m). In this case, the meta search still involves performing a binary search on four sections, resulting in a time complexity of O(log n). The subsequent binary search within the identified section has a time complexity of O(log m). Therefore, the overall time complexity for this example is O(log n + log m).

Example 4:
```go
sections := [][]string{
	{"Adi Parva", "Amba", "Bhishma", "Chitrangada"},
	{"Karna Parva", "Karna", "Kunti"},
	{"Virata Parva", "Draupadi", "Arjuna"},
	{"Udyoga Parva", "Yudhishthira", "Duryodhana"},
}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the target character is "Krishna", which is not found in the Mahabharata character list. The Meta Binary Search algorithm performs a meta search on the sections, but none of the sections contain the target character. Therefore, it concludes that the target character is not present in the Mahabharata.

The time complexity of the Meta Binary Search algorithm, in this case, is determined by the number of sections (n) and the number of elements in the largest section (m). The meta search involves performing a binary search on all four sections, resulting in a time complexity of O(log n). Since the target character is not found in any of the sections, there is no subsequent binary search within a specific section. Therefore, the overall time complexity for this example remains O(log n).

## Conclusion
In summary, the Meta Binary Search algorithm is applied in the code to search for a target character within specific sections of the Mahabharata. It involves performing a meta binary search to identify the section that may contain the target character, followed by a binary search within that section to find the exact index. The time complexity of the Meta Binary Search algorithm is O(log n) for the meta search and O(log m) for the binary search, where n is the number of sections and m is the number of characters in the identified section.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!