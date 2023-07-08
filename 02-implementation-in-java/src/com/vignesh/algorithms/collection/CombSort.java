package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * 
 * Comb Sort is an improvement over Bubble Sort that addresses the "turtles" problem in Bubble Sort, where small elements at the end of 
 * the array take a long time to reach their correct positions. By introducing the concept of a gap, Comb Sort performs comparisons 
 * and swaps over wider distances, gradually reducing the gap size and bringing larger elements closer to their correct positions. 
 * In this implementation, the combSort method performs the Comb Sort algorithm by comparing and swapping elements with a specific 
 * gap size until the gap becomes 1. The result is a fully sorted array of army sizes.
 *
 * Please note that the choice of a shrink factor of 1.3 and the specific gap reduction mechanism used in this implementation may vary 
 * in different versions of the Comb Sort algorithm.
 * 
 * 1. The combSort method takes an array of army sizes as input.
 * 2. It initializes the variable n with the length of the array.
 * 3. The sorting process begins by setting the initial gap size to the length of the array.
 * 4. A shrink factor of 1.3 is set to determine how much the gap is reduced in each iteration.
 * 5. The algorithm continues sorting until the gap becomes 1.
 * 6. Within each iteration, the algorithm calculates the current gap size by dividing the previous gap by the shrink factor and 
 *    converting it to an integer.
 * 7. The algorithm then traverses the array and compares elements with a specific gap.
 * 8. If the current element is greater than the element at a distance of the gap, the elements are swapped.
 * 9. This process continues for each element in the array, effectively reducing the number of inversions and bringing 
 *    larger elements closer to their correct positions.
 * 10. The gap size is reduced in each iteration, allowing the algorithm to focus on smaller elements.
 * 11. The process continues until the gap size becomes 1, at which point the algorithm performs a final pass with a 
 *     gap of 1 to ensure the array is fully sorted.
 * 12. Finally, the array of army sizes is fully sorted in ascending order.
 * 13. The sorted array of army sizes is printed to the console.
 *
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class CombSort {
  // Method to perform Comb Sort
  public static void combSort(int[] armySizes) {
    int n = armySizes.length;
    // Initialize the gap size
    int gap = n;
    // Set the shrink factor
    double shrink = 1.3;
    // Perform sorting until the gap becomes 1
    while (gap > 1) {
      // Calculate the current gap size
      gap = (int) (gap / shrink);
      // Traverse the array and swap elements with a specific gap
      for (int i = 0; i + gap < n; i++) {
        if (armySizes[i] > armySizes[i + gap]) {
          // Swap the elements
          int temp = armySizes[i];
          armySizes[i] = armySizes[i + gap];
          armySizes[i + gap] = temp;
        }
      }
    }
  }

  // Main method to test the Comb Sort implementation
  public static void main(String[] args) {
    int[] armySizes = { 10000, 5000, 2000, 7000, 1000, 3000, 6000, 8000, 4000 };
    System.out.println("Original army sizes: " + Arrays.toString(armySizes));
    // Call the Comb Sort method
    combSort(armySizes);
    System.out.println("Sorted army sizes: " + Arrays.toString(armySizes));
  }
}
