# Jump Search Algorithm in Go

This repository contains an implementation of the Jump Search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.


## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Explanation](#explanation)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The Jump Search algorithm is an efficient searching technique used to find a target element in a sorted array. It works by dividing the search range into blocks and performing a linear search within each block. The algorithm follows these steps:

1. Determine the block size (`step`) by taking the square root of the array length (`n`).
2. Find the block where the target element might be present by incrementing the `step` variable until the value at index `step` or the end of the array is greater than the target element.
3. Perform a linear search within the identified block, starting from the previous index (`prev`) until either the target element is found or the end of the block is reached.
4. If the target element is found, return its index. Otherwise, return -1 to indicate that the target element is not present in the array.

The time complexity of the Jump Search algorithm is O(sqrt(n)), where n is the number of elements in the array. It is faster than linear search but slower than binary search.

## Usage

The main program file (`main.go`) contains an example of how to use the Jump Search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Explanation
In the implemented code, the Jump Search algorithm is applied to search for a target character in the Mahabharata character list.
1. The `jumpSearch` function performs the Jump Search on a sorted slice of characters (`arr`) to find the target character. It returns the index of the target character if found or -1 if not found.
2. In the `main` function, a sorted slice called `characters` is created, representing the characters from the Mahabharata.
3. The `jumpSearch` function is called to find a character. If the character is found, its index is returned. Otherwise, -1 is returned.
Finally, the program prints whether the character was found or not, along with its index (if found).

## Examples
The program includes multiple examples to demonstrate the Jump search algorithm and its time complexity.

Example 1:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Draupadi"
```
Output : 
```shell
The character 'Draupadi' is found at index 4 in the Mahabharata.
```
In this example, the character "Draupadi" is found at index 4 using the Jump Search algorithm. The algorithm efficiently jumps ahead by block sizes to narrow down the search range, making it faster than linear search. The time complexity of Jump Search is O(sqrt(n)), where n is the number of elements in the array.

Example 2:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the character "Krishna" is not found in the Mahabharata character list. The Jump Search algorithm efficiently searches for the target character by jumping ahead in blocks, but since the character is not present, the algorithm determines its absence. The time complexity for this example remains O(sqrt(n)).

Example 3:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the target character is "Krishna", which is not found in the Mahabharata character list. The Jump Search algorithm divides the search range into blocks and performs a linear search within each block. Since the target character is not present, the algorithm determines its absence in the array.

## Conclusion
The time complexity of the Jump Search algorithm, in this case, is determined by the number of elements in the array (n). The algorithm divides the array into blocks, making jumps to narrow down the search range. Therefore, the time complexity of the Jump Search algorithm is O(sqrt(n)).

In summary, in Example 3, the Jump Search algorithm is applied to search for the character "Krishna" in the Mahabharata character list. As the character is not found, the algorithm efficiently determines its absence using the Jump Search technique. The time complexity for this example remains O(sqrt(n)), where n is the number of elements in the array.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!