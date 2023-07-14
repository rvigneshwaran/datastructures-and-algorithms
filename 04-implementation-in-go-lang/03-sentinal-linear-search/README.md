# Sentinel Linear Search Algorithm in Go

This repository contains an implementation of the Sentinel Linear Search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.

## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Explanation](#explanation)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The Sentinel Linear Search algorithm is an optimized version of the linear search algorithm. It works by adding a sentinel element to the end of the array being searched. The sentinel is chosen to be the target element itself, which simplifies the loop condition and allows for an early termination of the search loop when the target element is not found.

The time complexity of the binary search algorithm is O(log n), where n is the number of elements in the array. In each step, the algorithm reduces the search space by half, resulting in a faster search compared to linear search.

## Usage

The main program file (`main.go`) contains an example of how to use the sentinal linear search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Explanation
In the implemented code, the Sentinel Linear Search algorithm is applied to search for a target character in the Mahabharata character list.
1. The `searchCharacter` function takes in a slice of characters (`characters`) and the target character (`target`). It returns the index of the target character if found or -1 if not found.
2. The last element of the `characters` slice is temporarily stored as a sentinel using the `lastElement` variable.
3. The last element of the `characters` slice is replaced with the target character, enabling an early termination of the search loop if the target character is not found.
4. The search loop iterates through the array from the first element until the target character is found or the sentinel is encountered. Inside the loop, the index is incremented to move to the next element.
5. Once the search loop completes, the last element of the `characters` slice is restored to its original value using the `lastElement` variable.
6. The function checks if the target character was found based on the loop termination condition. If found, it returns the index of the target character. Otherwise, it returns -1.

## Examples
The program includes multiple examples to demonstrate the Sentinel linear search algorithm and its time complexity.

Example 1:
```go
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is found at index 6 in the Mahabharata.
```
In this example, the character "Krishna" is found at index 6. The sentinel linear search algorithm performs a linear search by iterating through the array until it finds a match or reaches the sentinel element. Since the target element is present in the array, the loop terminates early, resulting in improved performance compared to the regular linear search. Therefore, the time complexity is O(n), similar to the linear search.

Example 2:
```go
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Karna"
```
Output : 
```shell
The character 'Karna' is not found in the Mahabharata.
```
In this example, the character "Karna" is not found in the characters array. The sentinel linear search algorithm continues iterating through the array until it reaches the sentinel element without finding a match. The time complexity remains O(n), as the worst-case scenario still requires examining all elements.

Example 3:
```go
characters := []string{"Yudhishthira", "Bhima", "Arjuna", "Nakula", "Sahadeva", "Draupadi", "Krishna", "Duryodhana"}
character := "Yudhishthira"
```
Output : 
```shell
The character 'Yudhishthira' is found at index 0 in the Mahabharata.
```
In this example, the character "Yudhishthira" is found at the beginning of the characters array. The sentinel linear search algorithm compares each element until it finds a match. Since the target element is present at the first index, it terminates early, resulting in improved performance. The best-case time complexity is O(1), similar to the regular linear search.

## Conclusion
In summary, the sentinel linear search algorithm offers improved performance over the regular linear search by using a sentinel element. However, the time complexity remains O(n) in the worst case, where n is the number of elements in the array. 

The sentinel linear search is beneficial when the target element is more likely to be present near the beginning of the array, reducing the number of comparisons required.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!