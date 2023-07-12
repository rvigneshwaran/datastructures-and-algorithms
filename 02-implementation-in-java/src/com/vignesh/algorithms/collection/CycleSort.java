package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * 
 * Cycle Sort is an in-place comparison sorting algorithm known for its efficiency with minimal writes. It works by moving each element to 
 * its correct position, repeating the process until all elements are in their proper places. In this implementation, the cycleSort method 
 * performs the Cycle Sort algorithm by traversing the array and finding the correct position for each element. It then swaps the elements 
 * to place them in their proper positions. The result is a fully sorted array of army sizes.
 * 
 * Please note that Cycle Sort is particularly useful when the number of writes to memory is a concern, as it minimizes the number of 
 * writes by moving each element to its correct position. However, the number of writes is proportional to the number of elements 
 * that are out of order, so the efficiency of Cycle Sort can vary depending on the initial ordering of the elements.
 * 
 * 1.The cycleSort method takes an array of army sizes as input.
 * 2.It initializes the variable n with the length of the array.
 * 3.The sorting process begins by traversing the array from left to right.
 * 4.For each element at position i, the algorithm finds its correct position in the sorted array.
 * 5.The variable currentValue stores the value of the element being considered.
 * 6.The algorithm searches for the correct position by comparing the element with the remaining elements in the array.
 * 7.If an element is found to be smaller than currentValue, the correctPosition is incremented.
 * 8.Once the correct position is found, the algorithm checks if the element is already in the correct position.
 * 9.If it is, the algorithm proceeds to the next element in the array.
 * 10.If the element is not in the correct position, the algorithm moves all the elements greater than currentValue one position to the right.
 * 11.The element at the correct position is then swapped with currentValue, placing it in its correct position.
 * 12.The process continues until the element originally at position i is placed in its correct position.
 * 13.The algorithm then repeats the process for the next unsorted element until all elements are in their proper places.
 * 14.Finally, the array of army sizes is fully sorted in ascending order.
 * 15.The sorted array of army sizes is printed to the console.
 *
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class CycleSort {
  // Method to perform Cycle Sort
  public static void cycleSort(int[] armySizes) {
    int n = armySizes.length;
    // Traverse the array
    for (int i = 0; i < n - 1; i++) {
      int currentValue = armySizes[i];
      // Find the correct position for the current value
      int correctPosition = i;
      for (int j = i + 1; j < n; j++) {
        if (armySizes[j] < currentValue) {
          correctPosition++;
        }
      }
      // If the current value is already in the correct position, skip
      if (correctPosition == i) {
        continue;
      }
      // Move all the elements that are greater than the current value to the right
      while (currentValue == armySizes[correctPosition]) {
        correctPosition++;
      }
      // Swap the current value with the element at the correct position
      if (correctPosition != i) {
        int temp = armySizes[correctPosition];
        armySizes[correctPosition] = currentValue;
        currentValue = temp;
      }
      // Place the current value in its correct position by repeating the process
      while (correctPosition != i) {
        correctPosition = i;
        // Find the correct position for the current value
        for (int j = i + 1; j < n; j++) {
          if (armySizes[j] < currentValue) {
            correctPosition++;
          }
        }
        // Move all the elements that are equal to the current value to the right
        while (currentValue == armySizes[correctPosition]) {
          correctPosition++;
        }
        // Swap the current value with the element at the correct position
        if (currentValue != armySizes[correctPosition]) {
          int temp = armySizes[correctPosition];
          armySizes[correctPosition] = currentValue;
          currentValue = temp;
        }
      }
    }
  }

  // Main method to test the Cycle Sort implementation
  public static void main(String[] args) {
    int[] armySizes = { 10000, 5000, 2000, 7000, 1000, 3000, 6000, 8000, 4000 };
    System.out.println("Original army sizes: " + Arrays.toString(armySizes));
    // Call the Cycle Sort method
    cycleSort(armySizes);
    System.out.println("Sorted army sizes: " + Arrays.toString(armySizes));
  }
}
