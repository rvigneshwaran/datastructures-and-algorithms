# Ternary Search Algorithm in Go

This repository contains an implementation of the Ternary Search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.

## Algorithm

The Ternary Search algorithm is an efficient searching technique used to find a target element in a sorted array. It works by repeatedly dividing the search range into three equal parts and comparing the target element with the elements at two points within these parts. The algorithm follows these steps:

1. Initialize `left` and `right` variables to represent the boundaries of the search range within the sorted array.
2. Repeat until `left` becomes greater than `right`:
   a. Divide the search range into three equal parts using the indices `mid1 = left + (right-left)/3` and `mid2 = right - (right-left)/3`.
   b. Compare the target element with `arr[mid1]` and `arr[mid2]`.
   c. If the target element is found at `arr[mid1]`, return `mid1`.
   d. If the target element is found at `arr[mid2]`, return `mid2`.
   e. If the target element is less than `arr[mid1]`, update `right` to be `mid1 - 1`.
   f. If the target element is greater than `arr[mid2]`, update `left` to be `mid2 + 1`.
   g. If the target element is between `arr[mid1]` and `arr[mid2]`, update `left` to be `mid1 + 1` and `right` to be `mid2 - 1`.
3. If the target element is not found after the loop, return -1 to indicate its absence in the array.

The time complexity of the Ternary Search algorithm is O(log3 n), where n is the number of elements in the array. It is faster than linear search but slower than binary search due to the additional comparisons required to divide the search range into three parts.

## Usage

The main program file (`main.go`) contains an example of how to use the Ternary Search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Explanation
In the implemented code, the Ternary Search algorithm is applied to search for a target character in the Mahabharata character list.

1. The `ternarySearch` function performs the Ternary Search on a sorted slice of characters (`arr`) to find the target 
character. It returns the index of the target character if found or -1 if not found.
2. In the `main` function, a sorted slice called `characters` is created, representing the characters from the Mahabharata.
3. The `ternarySearch` function is called to find a character. If the character is found, its index is returned. Otherwise, -1 is returned.
4. Finally, the program prints whether the character was found or not, along with its index (if found).

## Examples
The program includes multiple examples to demonstrate the Ternary search algorithm and its time complexity.

Example 1:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Draupadi"
```
Output : 
```shell
The character 'Draupadi' is found at index 4 in the Mahabharata.
```
In this example, the character "Draupadi" is found at index 4. The Ternary Search algorithm efficiently narrows down the search range by dividing it into three parts, eliminating two-thirds of the remaining elements in each iteration. This results in faster search compared to linear search. Therefore, the time complexity of the Ternary Search algorithm is O(log3 n), where n is the number of elements in the array.

Example 2:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the character "Krishna" is not found in the Mahabharata character list. The Ternary Search algorithm eliminates two-thirds of the remaining elements in each iteration, efficiently searching for the target character. Since the target character is not present, the algorithm concludes its absence. The time complexity for this example remains O(log3 n).

Example 3:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the target character is "Krishna", which is not found in the Mahabharata character list. The Ternary Search algorithm divides the search range into three equal parts and compares the target character with the elements at two points within each part. Since the target character is not present, the algorithm concludes its absence in the array.

The time complexity of the Ternary Search algorithm, in this case, is determined by the number of elements in the array (n). The algorithm divides the search range into three parts, eliminating two-thirds of the remaining elements in each iteration. Therefore, the time complexity of the Ternary Search algorithm is O(log3 n).

In summary, in Example 3, the Ternary Search algorithm is applied to search for the character "Krishna" in the Mahabharata character list. As the character is not found, the algorithm determines its absence efficiently using the Ternary Search technique. The time complexity for this example remains O(log3 n), where n is the number of elements in the array.

## Conclusion
In summary, the Ternary Search algorithm efficiently searches for a target character in a sorted array by repeatedly dividing the search range into three equal parts. Its time complexity is O(log3 n), making it faster than linear search but slower than binary search.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!