// In this implementation, the functions bitonic_merge, compare_swap, and bitonic_sort are used to perform the 
// necessary steps of the Bitonic Sort algorithm. The bitonic_sort function is the main entry point for sorting the array.

// Function to perform a bitonic merge
fn bitonic_merge(arr: &mut [u32], start: usize, length: usize, direction: bool) {
    if length > 1 {
        let half = length / 2;
        for i in start..start + half {
            compare_swap(arr, i, i + half, direction);
        }
        bitonic_merge(arr, start, half, direction);
        bitonic_merge(arr, start + half, half, direction);
    }
}

// Function to perform a bitonic compare and swap
fn compare_swap(arr: &mut [u32], i: usize, j: usize, direction: bool) {
    if (arr[i] > arr[j] && direction) || (arr[i] < arr[j] && !direction) {
        arr.swap(i, j);
    }
}

// Function to perform a bitonic sort
fn bitonic_sort(arr: &mut [u32], start: usize, length: usize, direction: bool) {
    if length > 1 {
        let half = length / 2;
        bitonic_sort(arr, start, half, true);
        bitonic_sort(arr, start + half, half, false);
        bitonic_merge(arr, start, length, direction);
    }
}

// The test module contains test cases for the best-case scenario, worst-case scenario, 
// an empty array, an array with a single element, and an array with duplicate elements. 
// Each test case measures the execution time using Instant to provide an indication of 
// the algorithm's performance. The sorted arrays are then compared with the expected 
// results using assert_eq! to ensure the correctness of the sorting algorithm.

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::{Instant};

    #[test]
    fn test_bitonic_sort() {
        // Test case 1: Best-case scenario
        let mut arr1 = vec![1, 2, 3, 4, 5];
        let start1 = Instant::now();
        bitonic_sort(&mut arr1, 0, arr1.len(), true);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert_eq!(arr1, vec![1, 2, 3, 4, 5]);

        // Test case 2: Worst-case scenario
        let mut arr2 = vec![5, 4, 3, 2, 1];
        let start2 = Instant::now();
        bitonic_sort(&mut arr2, 0, arr2.len(), true);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert_eq!(arr2, vec![1, 2, 3, 4, 5]);

        // Test case 3: Empty array
        let mut arr3: Vec<u32> = vec![];
        let start3 = Instant::now();
        bitonic_sort(&mut arr3, 0, arr3.len(), true);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert_eq!(arr3, vec![]);

        // Test case 4: Array with a single element
        let mut arr4 = vec![42];
        let start4 = Instant::now();
        bitonic_sort(&mut arr4, 0, arr4.len(), true);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert_eq!(arr4, vec![42]);

        // Test case 5: Array with duplicate elements
        let mut arr5 = vec![7, 5, 9, 5, 3, 7];
        let start5 = Instant::now();
        bitonic_sort(&mut arr5, 0, arr5.len(), true);
        let elapsed5 = start5.elapsed();
        println!("Array with duplicate elements: {:?}", elapsed5);
        assert_eq!(arr5, vec![3, 5, 5, 7, 7, 9]);
    }
}
