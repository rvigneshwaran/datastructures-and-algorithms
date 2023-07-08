// Function to perform Cocktail Sort
fn cocktail_sort(arr: &mut [u32]) {
    let n = arr.len(); // Get the length of the array
    let mut start = 0; // Initialize the start index for forward pass
    let mut end = n - 1; // Initialize the end index for backward pass
    let mut swapped = true; // Flag to track if any elements were swapped

    while swapped {
        swapped = false; // Reset the swapped flag

        // Perform a forward pass
        for i in start..end {
            if arr[i] > arr[i + 1] {
                arr.swap(i, i + 1); // Swap elements if they are in the wrong order
                swapped = true; // Set the swapped flag to true
            }
        }

        if !swapped {
            break; // If no elements were swapped, the array is already sorted
        }

        swapped = false; // Reset the swapped flag

        // Perform a backward pass
        for i in (start..end).rev() {
            if arr[i] > arr[i + 1] {
                arr.swap(i, i + 1); // Swap elements if they are in the wrong order
                swapped = true; // Set the swapped flag to true
            }
        }

        start += 1; // Move the start index one position ahead
        end -= 1; // Move the end index one position back
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::{Duration, Instant};

    #[test]
    fn test_cocktail_sort() {
        // Test case 1: Best-case scenario (Array with elements in ascending order)
        let mut arr1 = vec![1, 2, 3, 4, 5];
        let start1 = Instant::now();
        cocktail_sort(&mut arr1);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert_eq!(arr1, vec![1, 2, 3, 4, 5]);

        // Test case 2: Worst-case scenario (Array with elements in descending order)
        let mut arr2 = vec![5, 4, 3, 2, 1];
        let start2 = Instant::now();
        cocktail_sort(&mut arr2);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert_eq!(arr2, vec![1, 2, 3, 4, 5]);

        // Test case 3: Empty array
        let mut arr3: Vec<u32> = vec![];
        let start3 = Instant::now();
        cocktail_sort(&mut arr3);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert_eq!(arr3, vec![]);

        // Test case 4: Array with a single element
        let mut arr4 = vec![42];
        let start4 = Instant::now();
        cocktail_sort(&mut arr4);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert_eq!(arr4, vec![42]);

        // Test case 5: Array with duplicate elements
        let mut arr5 = vec![7, 5, 9, 5, 3, 7];
        let start5 = Instant::now();
        cocktail_sort(&mut arr5);
        let elapsed5 = start5.elapsed();
        println!("Array with duplicate elements: {:?}", elapsed5);
        assert_eq!(arr5, vec![3, 5, 5, 7, 7, 9]);
    }
}
