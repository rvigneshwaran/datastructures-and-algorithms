package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * 
 * The Bubble Sort algorithm is implemented in the bubbleSort method. The algorithm works by repeatedly comparing adjacent 
 * elements and swapping them if they are in the wrong order. This process is repeated until the array is sorted.
 * 
 * The outer loop in bubbleSort performs n-1 passes, where n is the length of the array. The inner loop performs 
 * comparisons and swaps in each pass, moving the largest element to the end of the unsorted portion of the array.
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class BubbleSort {
  // Method to perform Bubble Sort
  public static void bubbleSort(int[] wandLengths) {
    // Get the length of the array
    int numWands = wandLengths.length;
    // Perform (numWands - 1) passes
    for (int i = 0; i < numWands - 1; i++) {
      // Perform comparisons and swaps in each pass
      for (int j = 0; j < numWands - i - 1; j++) {
        // Compare adjacent elements and swap if necessary
        if (wandLengths[j] > wandLengths[j + 1]) {
          int temp = wandLengths[j];
          wandLengths[j] = wandLengths[j + 1];
          wandLengths[j + 1] = temp;
        }
      }
    }
  }

  // Main method to test the Bubble Sort implementation
  public static void main(String[] args) {
    int[] wandLengths = { 9, 5, 2, 7, 1, 3, 6, 8, 4 };
    System.out.println("Original wand lengths: " + Arrays.toString(wandLengths));
    // Call the Bubble Sort method
    bubbleSort(wandLengths);
    System.out.println("Sorted wand lengths: " + Arrays.toString(wandLengths));
  }
}
