
fn bead_sort(mut arr: &mut [u32]) {
    // Find the maximum value in the array
    let max = *arr.iter().max().unwrap();
    // Get the length of the array
    let len = arr.len();

    // Create a grid to represent the beads
    let mut grid = vec![vec![0; len]; max as usize];

    // Place beads on the grid based on their values
    for i in 0..len {
        let mut beads = arr[i];
        for j in 0..max as usize {
            // Place beads on the grid row by row
            grid[j][i] = beads % 2;
            beads /= 2;
        }
    }

    // Move the beads down the grid, simulating gravity
    for i in 0..max as usize {
        let mut sum = 0;
        for j in 0..len {
            sum += grid[i][j];
            // Move beads down until the sum is exhausted
            grid[i][j] = 0;
            if sum > 0 {
                grid[i][j] = 1;
                sum -= 1;
            }
        }
    }

    // Calculate the sorted values based on the grid configuration
    for i in 0..len {
        let mut row_sum = 0;
        for j in 0..max as usize {
            row_sum += grid[j][i];
        }
        arr[i] = row_sum as u32;
    }
}

// The Big O notation for Bead Sort is as follows:

// Best-case scenario: O(n), where n is the number of elements in the array. In the best-case scenario, the array 
// is already sorted, so there is no need for any swapping or reordering.

// Worst-case scenario: O(n * m), where n is the number of elements in the array and m is the maximum value 
// in the array. In the worst-case scenario, the array is sorted in reverse order, so every bead needs to 
// fall to the bottom row before the sorting is complete.

// By measuring the execution time for each scenario, you can observe the actual time complexity of 
// the algorithm and compare it with the theoretical Big O notation.

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::{Instant};

    #[test]
    fn test_bead_sort() {
        // Test case 1: Best-case scenario
        let mut arr1 = vec![1, 2, 3, 4, 5];
        let start1 = Instant::now();
        bead_sort(&mut arr1);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert_eq!(arr1, vec![1, 2, 3, 4, 5]);

        // Test case 2: Worst-case scenario
        let mut arr2 = vec![5, 4, 3, 2, 1];
        let start2 = Instant::now();
        bead_sort(&mut arr2);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert_eq!(arr2, vec![1, 2, 3, 4, 5]);

        // Test case 3: Empty array
        let mut arr3: Vec<u32> = vec![];
        let start3 = Instant::now();
        bead_sort(&mut arr3);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert_eq!(arr3, vec![]);

        // Test case 4: Array with a single element
        let mut arr4 = vec![42];
        let start4 = Instant::now();
        bead_sort(&mut arr4);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert_eq!(arr4, vec![42]);

        // Test case 5: Array with duplicate elements
        let mut arr5 = vec![7, 5, 9, 5, 3, 7];
        let start5 = Instant::now();
        bead_sort(&mut arr5);
        let elapsed5 = start5.elapsed();
        println!("Array with duplicate elements: {:?}", elapsed5);
        assert_eq!(arr5, vec![3, 5, 5, 7, 7, 9]);
    }
}
