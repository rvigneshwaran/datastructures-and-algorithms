// bogo sort is a highly inefficient and unreliable algorithm that involves shuffling the elements 
// of an array randomly until the array is sorted. It is not practical for sorting large or 
// unsorted arrays and does not have a predictable time complexity.
// Note : The algorithm may not produce consistent results within a reasonable amount of time.

use rand::seq::SliceRandom;
use rand::thread_rng;

// Function to check if an array is sorted
fn is_sorted(arr: &[u32]) -> bool {
    // Iterate through the array and compare adjacent elements
    for i in 0..arr.len() - 1 {
        // If any element is greater than its next element, the array is not sorted
        if arr[i] > arr[i + 1] {
            return false;
        }
    }
    // If all elements are in non-decreasing order, the array is sorted
    true
}

// Function to shuffle the elements of an array randomly
fn shuffle(arr: &mut [u32]) {
    // Initialize the random number generator
    let mut rng = thread_rng();
    // Shuffle the elements using the Fisher-Yates algorithm
    arr.shuffle(&mut rng);
}

// Function to perform bogo sort
fn bogo_sort(arr: &mut [u32]) {
    // Repeat until the array is sorted
    while !is_sorted(arr) {
        // Shuffle the array randomly
        shuffle(arr);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::{Duration, Instant};

    #[test]
    fn test_bogo_sort() {
        // Test case 1: Best-case scenario
        let mut arr1 = vec![1, 2, 3, 4, 5];
        let start1 = Instant::now();
        bogo_sort(&mut arr1);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert!(is_sorted(&arr1));

        // Test case 2: Worst-case scenario (Array with descending order)
        let mut arr2 = vec![5, 4, 3, 2, 1];
        let start2 = Instant::now();
        bogo_sort(&mut arr2);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert!(is_sorted(&arr2));

        // Test case 3: Empty array
        let mut arr3: Vec<u32> = vec![];
        let start3 = Instant::now();
        bogo_sort(&mut arr3);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert!(is_sorted(&arr3));

        // Test case 4: Array with a single element
        let mut arr4 = vec![42];
        let start4 = Instant::now();
        bogo_sort(&mut arr4);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert!(is_sorted(&arr4));
    }
}
