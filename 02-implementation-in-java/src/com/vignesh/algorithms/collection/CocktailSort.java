package com.vignesh.algorithms.collection;

/**
 * 
 * Cocktail Sort, also known as Bidirectional Bubble Sort, is a variation of the Bubble Sort algorithm that sorts in both directions. 
 * By performing multiple passes of traversing the array from left to right and then from right to left, the largest and smallest 
 * elements gradually move to their correct positions. In this implementation, the cocktailSort method performs the 
 * Cocktail Sort algorithm by repeatedly traversing the array in both directions and swapping adjacent elements as needed. 
 * The result is a fully sorted array of army sizes.
 * 
 * Please note that Cocktail Sort is an optimized version of Bubble Sort that reduces the number of swaps required. 
 * It performs slightly better than Bubble Sort but is still less efficient than more advanced sorting algorithms for larger data sets.
 * 
 * 1.The cocktailSort method takes an array of army sizes as input.
 * 2.It initializes the variable n with the length of the array.
 * 3.The sorting process begins by setting the starting values for the left and right pointers.
 * 4.The left pointer is set to the leftmost position (0) of the array.
 * 5.The right pointer is set to the rightmost position (n-1) of the array.
 * 6.The algorithm continues sorting until the left pointer surpasses the right pointer.
 * 7.Within each iteration, the algorithm traverses the array from left to right, moving the largest element to the rightmost position.
 * 8.The for loop is used to compare adjacent elements and swap them if necessary.
 * 9.If the current element is greater than the next element, they are swapped.
 * 10.After each traversal, the right pointer is decremented, as the largest element is now in its correct position at the 
 *    rightmost position of the remaining unsorted portion of the array.
 * 11.The algorithm then traverses the array from right to left, moving the smallest element to the leftmost position.
 * 12.The for loop is used to compare adjacent elements and swap them if necessary.
 * 13.If the current element is smaller than the previous element, they are swapped.
 * 14.After each traversal, the left pointer is incremented, as the smallest element is now in its correct position 
 *    at the leftmost position of the remaining unsorted portion of the array.
 * 15.The process continues until the left pointer surpasses the right pointer, indicating that the entire array is sorted.
 * 16.Finally, the array of army sizes is fully sorted in ascending order.
 * 17.The sorted array of army sizes is printed to the console.
 * 
 * 
 * @author Vigneshwaran Rethenasekaran
 *
 */
import java.util.Arrays;

public class CocktailSort {
  // Method to perform Cocktail Sort
  public static void cocktailSort(int[] armySizes) {
    int n = armySizes.length;
    // Set the starting values for the left and right pointers
    int left = 0;
    int right = n - 1;
    while (left < right) {
      // Traverse from left to right, moving the largest element to the rightmost position
      for (int i = left; i < right; i++) {
        if (armySizes[i] > armySizes[i + 1]) {
          // Swap the elements
          int temp = armySizes[i];
          armySizes[i] = armySizes[i + 1];
          armySizes[i + 1] = temp;
        }
      }
      right--;
      // Traverse from right to left, moving the smallest element to the leftmost position
      for (int i = right; i > left; i--) {
        if (armySizes[i] < armySizes[i - 1]) {
          // Swap the elements
          int temp = armySizes[i];
          armySizes[i] = armySizes[i - 1];
          armySizes[i - 1] = temp;
        }
      }
      left++;
    }
  }

  // Main method to test the Cocktail Sort implementation
  public static void main(String[] args) {
    int[] armySizes = { 10000, 5000, 2000, 7000, 1000, 3000, 6000, 8000, 4000 };
    System.out.println("Original army sizes: " + Arrays.toString(armySizes));
    // Call the Cocktail Sort method
    cocktailSort(armySizes);
    System.out.println("Sorted army sizes: " + Arrays.toString(armySizes));
  }
}
