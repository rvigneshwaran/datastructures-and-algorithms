package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * 
 * Counting Sort is a linear time sorting algorithm that works by counting the number of occurrences of each element in the 
 * array and then placing each element in its correct position based on the count array. In this implementation, the countingSort 
 * method performs the Counting Sort algorithm by counting the frequency of each element in the array, modifying the count array 
 * to represent the starting position of each element in the sorted array, and finally building the sorted array by placing each 
 * element in its correct position. This results in a fully sorted array of army sizes.
 * 
 * 1. The countingSort method takes an array of army sizes as input.
 * 2. It initializes the variable n with the length of the array.
 * 3. The sorting process begins by finding the maximum element in the array, which determines the size of the count array.
 * 4. The getMaxElement helper method is called to find the maximum element in the array.
 * 5. The count array is created with a size of max + 1, where max is the maximum element in the array.
 * 6. The count array is used to store the frequency of each element in the array.
 * 7. A loop traverses the array, incrementing the count for each element encountered.
 * 8. The count array is modified to store the actual position of each element in the sorted array.
 * 9. Another loop starts from index 1 up to max, and each element in the count array is updated by adding the previous count.
 * 10. This modification ensures that each element represents the starting position of its corresponding value in the sorted array.
 * 11. A temporary array called sortedArray is created to store the sorted elements.
 * 12. A loop starts from the end of the original array and iterates backwards.
 * 13. For each element in the original array, the corresponding position is retrieved from the count array and decremented.
 * 14. The element is then placed in its correct position in the sortedArray by accessing the calculated position.
 * 15. This process continues until all elements in the original array have been placed in the sortedArray.
 * 16. Finally, the sortedArray is copied back to the original array, resulting in a fully sorted array of army sizes.
 * 17. The sorted array of army sizes is printed to the console.
 * 
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class CountingSort {
  // Method to perform Counting Sort
  public static void countingSort(int[] armySizes) {
    int n = armySizes.length;
    // Find the maximum element in the array to determine the count array size
    int max = getMaxElement(armySizes);
    // Create a count array to store the frequency of each element
    int[] count = new int[max + 1];
    // Calculate the frequency of each element
    for (int i = 0; i < n; i++) {
      count[armySizes[i]]++;
    }
    // Modify the count array to store the actual position of each element in the sorted
    // array
    for (int i = 1; i <= max; i++) {
      count[i] += count[i - 1];
    }
    // Create a temporary array to store the sorted elements
    int[] sortedArray = new int[n];
    // Build the sorted array by placing each element in its correct position
    for (int i = n - 1; i >= 0; i--) {
      sortedArray[count[armySizes[i]] - 1] = armySizes[i];
      count[armySizes[i]]--;
    }
    // Copy the sorted array back to the original array
    System.arraycopy(sortedArray, 0, armySizes, 0, n);
  }

  // Helper method to get the maximum element in the array
  private static int getMaxElement(int[] armySizes) {
    int max = armySizes[0];
    for (int i = 1; i < armySizes.length; i++) {
      if (armySizes[i] > max) {
        max = armySizes[i];
      }
    }
    return max;
  }

  // Main method to test the Counting Sort implementation
  public static void main(String[] args) {
    int[] armySizes = { 10000, 5000, 2000, 7000, 1000, 3000, 6000, 8000, 4000 };
    System.out.println("Original army sizes: " + Arrays.toString(armySizes));
    // Call the Counting Sort method
    countingSort(armySizes);
    System.out.println("Sorted army sizes: " + Arrays.toString(armySizes));
  }
}
