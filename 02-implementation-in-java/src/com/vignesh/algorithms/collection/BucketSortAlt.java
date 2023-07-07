package com.vignesh.algorithms.collection;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

/**
 * Bucket Sort is a sorting algorithm that divides the input into several buckets and sorts the elements within each bucket individually. 
 * In this implementation, the wand powers are divided into buckets based on their range, and each bucket is sorted separately. 
 * The sorted elements from all the buckets are then combined to produce the final sorted array.
 * 
 * 1. The bucketSort method takes an array of wand powers as input.
 * 2. The maximum power and minimum power of the wand powers are determined using the getMaxPower and getMinPower helper methods.
 * 3. The range of wand powers is calculated by subtracting the minimum power from the maximum power and adding 1.
 * 4. An ArrayList of buckets is created, with each bucket representing a range of wand powers.
 * 5. Each bucket is initialized as an empty ArrayList.
 * 6. The wand powers are distributed into the buckets based on their values. Each wand power is subtracted by the 
 *    minimum power and used as an index to determine which bucket it belongs to.
 * 7. Each bucket is sorted individually using the Collections.sort method.
 * 8. The sorted elements from all the buckets are concatenated back into the original array.
 * 9. Finally, the sorted array of wand powers is returned.
 * 
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class BucketSortAlt {
  // Method to perform Bucket Sort
  public static void bucketSort(int[] wandPowers) {
    // Determine the range of wand powers
    int maxPower = getMaxPower(wandPowers);
    int minPower = getMinPower(wandPowers);
    // Calculate the number of buckets needed
    int numBuckets = maxPower - minPower + 1;
    // Create an array of buckets
    ArrayList<ArrayList<Integer>> buckets = new ArrayList<>(numBuckets);
    for (int indexI = 0; indexI < numBuckets; indexI++) {
      buckets.add(new ArrayList<>());
    }
    // Distribute wand powers into buckets
    for (int power : wandPowers) {
      int index = power - minPower;
      buckets.get(index).add(power);
    }
    // Sort each bucket
    for (ArrayList<Integer> bucket : buckets) {
      Collections.sort(bucket);
    }
    // Concatenate the sorted buckets back into the original array
    int index = 0;
    for (ArrayList<Integer> bucket : buckets) {
      for (int power : bucket) {
        wandPowers[index++] = power;
      }
    }
  }

  // Helper method to get the maximum power in the array
  private static int getMaxPower(int[] wandPowers) {
    int max = Integer.MIN_VALUE;
    for (int power : wandPowers) {
      if (power > max) {
        max = power;
      }
    }
    return max;
  }

  // Helper method to get the minimum power in the array
  private static int getMinPower(int[] wandPowers) {
    int min = Integer.MAX_VALUE;
    for (int power : wandPowers) {
      if (power < min) {
        min = power;
      }
    }
    return min;
  }

  // Main method to test the Bucket Sort implementation
  public static void main(String[] args) {
    int[] wandPowers = { 9, 35, 12, 7, 25, 17, 42, 28, 39 };
    System.out.println("Original wand powers: " + Arrays.toString(wandPowers));
    // Call the Bucket Sort method
    bucketSort(wandPowers);
    System.out.println("Sorted wand powers: " + Arrays.toString(wandPowers));
  }
}
