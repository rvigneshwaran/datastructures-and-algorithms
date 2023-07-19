# Interpolation Search Algorithm in Go (Mahabharata Characters)

This repository contains an implementation of the Interpolation Search algorithm in Go programming language, with the context of searching for a character in the epic Mahabharata.


## Table of Contends

- [Algorithm](#algorithm)
- [Usage](#usage)
- [Explanation](#explanation)
- [Examples](#examples)
- [Conclusion](#conclusion)

## Algorithm

The Interpolation Search algorithm is an efficient searching technique used to find a target element in a sorted array. It works by estimating the probable position of the target element within the search range using interpolation. The algorithm follows these steps:

1. Initialize the low index (`low`) as 0 and the high index (`high`) as the length of the array minus 1.
2. Repeat until the target element is within the range of `arr[low]` and `arr[high]`:
   - If `low` is equal to `high`, check if the element at that index is the target element. If so, return the index. Otherwise, the target element is not found, so return -1.
   - Calculate the probable position (`position`) of the target element using the interpolation formula:
     `position = low + floor((target - arr[low]) * (high - low) / (arr[high] - arr[low]))`
   - If the element at position `position` is the target element, return `position`.
   - If the element at position `position` is less than the target element, update `low` to be `position + 1`.
   - If the element at position `position` is greater than the target element, update `high` to be `position - 1`.
3. If the target element is not found after the loop, return -1 to indicate its absence in the array.

The time complexity of the Interpolation Search algorithm is O(log(log n)), where n is the number of elements in the array. It is more efficient than binary search in certain scenarios.

## Usage

The main program file (`main.go`) contains an example of how to use the Interpolation Search algorithm to search for characters in the Mahabharata.

To run the program:

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command:

```shell
go run main.go
```

## Explanation
1. The `interpolationSearch` function takes two parameters: the sorted slice of characters (`arr`) and the target character (`target`). It returns the index of the target character if found or -1 if not found.

2. Inside the `interpolationSearch` function, two variables, `low `and `high`, are initialized to represent the low and high indices of the search range within the array.

3. The algorithm uses a while loop to iterate until the target character is within the range of `arr[low]` and `arr[high]`. This ensures that the search range is valid and the target character may potentially be present.

4. Within the loop, the algorithm first checks if `low` is equal to `high`. If they are equal, it means there is only one element left in the search range. The algorithm checks if the element at that index is the target character. If it is, the function returns the index. If not, the target character is not found, so the function returns -1.

5. If `low` and `high` are not equal, the algorithm proceeds to calculate the probable position of the target character within the search range using the interpolation formula:
```scss
position = low + floor((target - arr[low]) * (high - low) / (arr[high] - arr[low]))
```
This formula estimates the position by considering the relative difference between the target character and the elements at the low and high indices. It uses linear interpolation to estimate the position.

6. After calculating the probable position (position), the algorithm checks if the element at that position is the target character. If it is, the function returns the position as the index of the target character.

7. If the element at the `position` is less than the target character, it means the target character may be located in the right half of the search range. The low index is updated to `position + 1` to search in the right half.

8. If the element at the `position` is greater than the target character, it means the target character may be located in the left half of the search range. The `high` index is updated to `position` - 1 to search in the left half.

9. The loop continues until the target character is found within the search range, or the range becomes invalid (`low` exceeds `high`), indicating that the target character is not present in the array.

10. If the target character is not found after the loop, the function returns -1 to indicate its absence in the array.

In the `main` function, a sorted slice called `characterList` is created to represent the characters from the Mahabharata. The interpolationSearch function is then called with the `characterList` and a `target` character, and the returned index is stored in the index variable.

Finally, the program checks if the index is not equal to -1. If it is not -1, it means the target character is found in the Mahabharata, and the program prints a message indicating the character and its index. If the index is -1, it means the target character is not found, and the program prints a message indicating the absence of the character.

That's how the Interpolation Search algorithm is applied in the above snippet to search for a target character in the Mahabharata character list. The algorithm estimates the probable position of the target character using interpolation and performs efficient searches in the sorted array.

## Examples
The program includes multiple examples to demonstrate the Interpolation Search algorithm and its time complexity.
Example 1:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Draupadi"
```
Output : 
```shell
The character 'Draupadi' is found at index 4 in the Mahabharata.
```
In this example, the character "Draupadi" is found at index 4 using the Interpolation Search algorithm. The algorithm uses interpolation to estimate the probable position of the target character within the search range, resulting in faster search compared to linear search. Therefore, the time complexity of the Interpolation Search algorithm is O(log(log n)), making it more efficient than binary search in certain scenarios.

Example 2:
```go
characters := []string{"Amba", "Arjuna", "Bhishma", "Chitrangada", "Draupadi", "Duryodhana", "Karna", "Kunti", "Yudhishthira"}
character := "Krishna"
```
Output : 
```shell
The character 'Krishna' is not found in the Mahabharata.
```
In this example, the character "Krishna" is not found in the Mahabharata character list. The Interpolation Search algorithm uses interpolation to estimate the position of the target character, but since the character is not present, the algorithm determines its absence. The time complexity for this example remains O(log(log n)).

## Conclusion
The Interpolation Search algorithm estimates the probable position of the target character using interpolation and performs efficient searches in the sorted array.

In summary, the Interpolation Search algorithm efficiently searches for a target character in a sorted array by estimating its position using interpolation. Its time complexity is O(log(log n)), making it faster than binary search in certain scenarios.

Feel free to modify and use the provided code as a starting point for your own projects.

If you have any questions or suggestions, feel free to reach out!