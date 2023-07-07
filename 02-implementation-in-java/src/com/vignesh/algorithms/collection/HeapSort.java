package com.vignesh.algorithms.collection;

import java.util.Arrays;

/**
 * Heap Sort is based on the concept of a binary heap, which is a complete binary tree that satisfies the heap property. 
 * In this algorithm, the array is treated as a binary heap, and the heapify operation is used to build the heap and 
 * maintain the heap property during the sorting process. By repeatedly extracting the maximum element from the heap 
 * and heapifying the remaining elements, the array is sorted in ascending order.
 * 
 * 1. The heapSort method takes an array of wand powers as input.
 * 2. It starts by initializing the length variable to the length of the array.
 * 3. The algorithm then builds a max heap from the array using the buildMaxHeap method.
 * 4. After building the max heap, the algorithm performs the sorting process by extracting elements from the heap one by one.
 * 5. In each iteration, the root (maximum element) of the heap is swapped with the last element in the unsorted portion of the array.
 * 6. The size of the heap is reduced by 1, effectively moving the largest element to its correct position at the end of the array.
 * 7. The algorithm then performs heapify on the reduced heap to restore the heap property.
 * 8. Steps 5-7 are repeated until the entire array is sorted in ascending order.
 * 9. Finally, the sorted array of wand powers is printed to the console.
 * 
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class HeapSort {
  // Method to perform Heap Sort
  public static void heapSort(int[] wandPowers) {
    int length = wandPowers.length;
    // Build the max heap
    buildMaxHeap(wandPowers, length);
    // Extract elements from the heap one by one
    for (int indexI = length - 1; indexI >= 0; indexI--) {
      // Swap the root (maximum element) with the last element
      int tempElement = wandPowers[0];
      wandPowers[0] = wandPowers[indexI];
      wandPowers[indexI] = tempElement;
      // Perform heapify on the reduced heap
      heapify(wandPowers, indexI, 0);
    }
  }

  // Method to build the max heap
  private static void buildMaxHeap(int[] wandPowers, int length) {
    // Build the heap from the second last level
    for (int indexI = length / 2 - 1; indexI >= 0; indexI--) {
      heapify(wandPowers, length, indexI);
    }
  }

  // Method to heapify a subtree rooted at index i
  private static void heapify(int[] wandPowers, int length, int indexI) {
    int largestELement = indexI; // Initialize the largest as root
    int leftChild = 2 * indexI + 1; // Left child
    int rightChild = 2 * indexI + 2; // Right child
    // If left child is larger than root
    if (leftChild < length && wandPowers[leftChild] > wandPowers[largestELement]) {
      largestELement = leftChild;
    }
    // If right child is larger than current largest
    if (rightChild < length && wandPowers[rightChild] > wandPowers[largestELement]) {
      largestELement = rightChild;
    }
    // If the largest is not the root
    if (largestELement != indexI) {
      // Swap the largest element with the root
      int temp = wandPowers[indexI];
      wandPowers[indexI] = wandPowers[largestELement];
      wandPowers[largestELement] = temp;
      // Recursively heapify the affected subtree
      heapify(wandPowers, length, largestELement);
    }
  }

  // Main method to test the Heap Sort implementation
  public static void main(String[] args) {
    int[] wandPowers = { 9, 5, 2, 7, 1, 3, 6, 8, 4 };
    System.out.println("Original wand powers: " + Arrays.toString(wandPowers));
    // Call the Heap Sort method
    heapSort(wandPowers);
    System.out.println("Sorted wand powers: " + Arrays.toString(wandPowers));
  }
}
