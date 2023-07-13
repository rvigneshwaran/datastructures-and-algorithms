# Binary Search Algorithm in Go (Mahabharata Characters)

This repository contains an implementation of the binary search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.

## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The binary search algorithm is a divide-and-conquer search algorithm used to find a specific element in a sorted array. It repeatedly divides the array in half and compares the middle element with the target element. Based on the comparison, it eliminates the half of the array where the target element cannot be present. It continues this process until a match is found or the subarray becomes empty.

The time complexity of the binary search algorithm is O(log n), where n is the number of elements in the array. In each step, the algorithm reduces the search space by half, resulting in a faster search compared to linear search.

## Usage

The main program file (`main.go`) contains an example of how to use the binary search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Examples
The program includes multiple examples to demonstrate the linear search algorithm and its time complexity.

Example 1:
```go
characters := []string{"Arjuna", "Bhima", "Draupadi", "Duryodhana", "Krishna", "Nakula", "Sahadeva", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is found at index 4 in the Mahabharata.
```
In this example, the character "Krishna" is found at index 4. The binary search algorithm divides the array into halves and compares the middle element with the target element. It continues dividing the subarray until it finds a match. Therefore, the time complexity is O(log n), where n is the number of elements in the array.

Example 2:
```go
characters := []string{"Arjuna", "Bhima", "Draupadi", "Duryodhana", "Krishna", "Nakula", "Sahadeva", "Yudhishthira"}
character := "Karna"
```
Output : 
```shell
The character 'Karna' is not found in the Mahabharata.
```
In this example, the character "Karna" is not found in the characters array. The binary search algorithm continues dividing the array into halves until it finds the target element or the subarray becomes empty. Since "Karna" is not present in the array, it returns -1. The worst-case time complexity remains O(log n).

Example 3:
```go
characters := []string{"Arjuna", "Bhima", "Draupadi", "Duryodhana", "Krishna", "Nakula", "Sahadeva", "Yudhishthira"}
character := "Arjuna"
```
Output : 
```shell
The character 'Arjuna' is found at index 0 in the Mahabharata.
```
In this example, the character "Arjuna" is found at the beginning of the characters array. The binary search algorithm quickly identifies the target element in the first comparison. Therefore, the best-case time complexity is O(1), which occurs when the desired element is found at the beginning.

## Conclusion
In summary, the binary search algorithm has a time complexity of O(log n) in the worst case, as it divides the array in half repeatedly to search for the target element. The best-case time complexity is O(1) when the desired element is found at the beginning. The binary search algorithm is more efficient than linear search for larger arrays due to its logarithmic time complexity.

The binary search algorithm is an efficient search algorithm for finding an element in a sorted array. Its time complexity of O(log n) allows for faster searching compared to linear search, especially for larger arrays.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!