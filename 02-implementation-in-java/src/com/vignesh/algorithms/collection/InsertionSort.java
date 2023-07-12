package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * The Insertion Sort algorithm works by dividing the array into two portions: a sorted portion on the left and 
 * an unsorted portion on the right. In each iteration, an element from the unsorted portion is selected, and 
 * it is compared and inserted into its correct position within the sorted portion. This process is repeated 
 * until the entire array is sorted.
 * 
 * 1. The insertionSort method takes an array of wand powers as input.
 * 2. It starts by iterating over the array, starting from the second element (index 1) to the last element 
 *    (index n-1), where n is the length of the array.
 * 3. For each element at index i, the algorithm assigns it to the variable key and sets indexJ as indexI - 1.
 * 4. The algorithm then compares the key with the element at index indexJ (one position before indexI) and 
 *    continues moving elements greater than the key to one position ahead (i.e., shifting them to the right) 
 *    until it finds the correct position for the key or reaches the beginning of the array.
 * 5. Once the correct position is found, the key is inserted at that position.
 * 6. The process continues for all elements in the array, resulting in a sorted array in ascending order.
 * 7. Finally, the sorted array of wand powers is printed to the console.
 * 
 * @author Vicgneshwaran Rethenasekaran
 *
 */
public class InsertionSort {
  // Method to perform Insertion Sort
  public static void insertionSort(int[] wandPowers) {
    int total_count = wandPowers.length;
    // Iterate over the array starting from the second element
    for (int indexI = 1; indexI < total_count; indexI++) {
      int individual_key = wandPowers[indexI];
      int indexJ = indexI - 1;
      // Move elements greater than the key to one position ahead
      while (indexJ >= 0 && wandPowers[indexJ] > individual_key) {
        wandPowers[indexJ + 1] = wandPowers[indexJ];
        indexJ--;
      }
      wandPowers[indexJ + 1] = individual_key;
    }
  }

  // Main method to test the Insertion Sort implementation
  public static void main(String[] args) {
    int[] wandPowers = { 9, 5, 2, 7, 1, 3, 6, 8, 4 };
    System.out.println("Original wand powers: " + Arrays.toString(wandPowers));
    // Call the Insertion Sort method
    insertionSort(wandPowers);
    System.out.println("Sorted wand powers: " + Arrays.toString(wandPowers));
  }
}
