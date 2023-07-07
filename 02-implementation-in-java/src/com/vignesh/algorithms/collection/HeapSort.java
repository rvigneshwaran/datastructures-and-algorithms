package com.vignesh.algorithms.collection;

import java.util.Arrays;

public class HeapSort {
  public static void heapSort(int[] wandPowers) {
    int length = wandPowers.length;
    buildMaxHeap(wandPowers, length);
    for (int indexI = length - 1; indexI >= 0; indexI--) {
      int tempElement = wandPowers[0];
      wandPowers[0] = wandPowers[indexI];
      wandPowers[indexI] = tempElement;
      heapify(wandPowers, indexI, 0);
    }
  }

  private static void buildMaxHeap(int[] wandPowers, int length) {
    for (int indexI = length / 2 - 1; indexI >= 0; indexI--) {
      heapify(wandPowers, length, indexI);
    }
  }

  private static void heapify(int[] wandPowers, int length, int indexI) {
    int largestELement = indexI; // Initialize the largest as root
    int leftChild = 2 * indexI + 1; // Left child
    int rightChild = 2 * indexI + 2; // Right child
    if (leftChild < length && wandPowers[leftChild] > wandPowers[largestELement]) {
      largestELement = leftChild;
    }
    if (rightChild < length && wandPowers[rightChild] > wandPowers[largestELement]) {
      largestELement = rightChild;
    }
    if (largestELement != indexI) {
      int temp = wandPowers[indexI];
      wandPowers[indexI] = wandPowers[largestELement];
      wandPowers[largestELement] = temp;
      heapify(wandPowers, length, largestELement);
    }
  }

  public static void main(String[] args) {
    int[] wandPowers = { 9, 5, 2, 7, 1, 3, 6, 8, 4 };
    System.out.println("Original wand powers: " + Arrays.toString(wandPowers));
    heapSort(wandPowers);
    System.out.println("Sorted wand powers: " + Arrays.toString(wandPowers));
  }
}