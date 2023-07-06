package com.vignesh.algorithms.collection;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

/**
 * The buckets represent different categories of wand lengths, and the sorting process resembles organizing wands 
 * based on their lengths. You can modify the getBucketIndex method to match the wand length categories or define 
 * your own logic based on the context you wish to represent.
 * 
 * @author Vigneshwaran Rethenasekaran
 *
 */
public class BucketSort {
  // Method to perform Bucket Sort
  public static void bucketSort(int[] wandLengths) {
    // Create a list of buckets (lists) for wand lengths
    ArrayList<ArrayList<Integer>> buckets = new ArrayList<>();
    // Initialize the buckets
    for (int i = 0; i < wandLengths.length; i++) {
      buckets.add(new ArrayList<>());
    }
    // Distribute the wand lengths into buckets
    for (int i = 0; i < wandLengths.length; i++) {
      int index = getBucketIndex(wandLengths[i]);
      buckets.get(index).add(wandLengths[i]);
    }
    // Sort each bucket
    for (ArrayList<Integer> bucket : buckets) {
      Collections.sort(bucket);
    }
    // Combine the sorted buckets into a single sorted array
    int index = 0;
    for (ArrayList<Integer> bucket : buckets) {
      for (int length : bucket) {
        wandLengths[index++] = length;
      }
    }
  }

  // Method to get the index of the bucket for a given wand length
  private static int getBucketIndex(int wandLength) {
    // Assume each bucket represents a specific wand length
    // category.
    // You can define your own logic here based on the wand length categories in the Harry
    // Potter universe.
    // For example, if wand lengths are categorized into different sizes (short, medium,
    // long),
    // you can use conditional statements to map wand lengths to bucket indices accordingly.
    // In this example, let's assume the buckets are represented by wand lengths 0-9, 10-19,
    // 20-29, etc.
    return wandLength / 10;
  }

  // Main method to test the Bucket Sort implementation
  public static void main(String[] args) {
    int[] wandLengths = { 9, 35, 12, 7, 25, 17, 42, 28, 39 };
    System.out.println("Original wand lengths: " + Arrays.toString(wandLengths));
    // Call the Bucket Sort method
    bucketSort(wandLengths);
    System.out.println("Sorted wand lengths: " + Arrays.toString(wandLengths));
  }
}
