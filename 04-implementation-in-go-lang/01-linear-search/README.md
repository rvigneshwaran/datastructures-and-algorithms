# Linear Search Algorithm in Go (Mahabharata Characters)

This repository contains an implementation of the linear search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.

## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The linear search algorithm performs a sequential search through an array to find a specific element. It iterates through each element of the array and compares it with the target element until a match is found or the end of the array is reached.

The time complexity of the linear search algorithm is O(n), where n is the number of elements in the array. In the worst case, the algorithm may need to iterate through all the elements to find the target element.

## Usage

The main program file (`main.go`) contains an example of how to use the linear search algorithm to search for characters in the Mahabharata.

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
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is found at index 6 in the Mahabharata.
```
In this example, the character "Krishna" is found at index 6. The linear search algorithm iterates through the characters array until it finds a match, which happens in this case. Therefore, the time complexity is O(n), where n is the number of elements in the array.

Example 2:
```go
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Karna"
```
Output : 
```shell
The character 'Karna' is not found in the Mahabharata.
```
In this example, the character "Karna" is not found in the characters array. The linear search algorithm iterates through the entire array without finding a match. As a result, it has to go through all the elements, giving a worst-case time complexity of O(n).

Example 3:
```go
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Yudhishthira"
```
Output : 
```shell
The character 'Yudhishthira' is found at index 0 in the Mahabharata.
```
In this example, the character "Yudhishthira" is found at the beginning of the characters array. The linear search algorithm checks the first element and finds a match. Therefore, the best-case time complexity is O(1), which occurs when the desired element is found at the beginning.

## Conclusion
The linear search algorithm is a simple and straightforward method for searching for an element in an array. However, its time complexity of O(n) makes it less efficient for larger arrays compared to other search algorithms like binary search.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!