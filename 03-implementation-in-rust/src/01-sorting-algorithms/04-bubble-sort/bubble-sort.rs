// Function to perform Bubble Sort
fn bubble_sort(arr: &mut [u32]) {
    let n = arr.len(); // Get the length of the array
    for i in 0..n { // Iterate through the array from the beginning
        for j in 0..n - i - 1 { // Iterate through the unsorted part of the array
            if arr[j] > arr[j + 1] { // Compare adjacent elements
                arr.swap(j, j + 1); // Swap elements if they are in the wrong order
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::Instant;

    #[test]
    fn test_bubble_sort() {
        // Test case 1: Best-case scenario (Already sorted array)
        let mut arr1 = vec![1, 2, 3, 4, 5];
        let start1 = Instant::now();
        bubble_sort(&mut arr1);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert_eq!(arr1, vec![1, 2, 3, 4, 5]);

        // Test case 2: Worst-case scenario (Array sorted in reverse order)
        let mut arr2 = vec![5, 4, 3, 2, 1];
        let start2 = Instant::now();
        bubble_sort(&mut arr2);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert_eq!(arr2, vec![1, 2, 3, 4, 5]);

        // Test case 3: Empty array
        let mut arr3: Vec<u32> = vec![];
        let start3 = Instant::now();
        bubble_sort(&mut arr3);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert_eq!(arr3, vec![]);

        // Test case 4: Array with a single element
        let mut arr4 = vec![42];
        let start4 = Instant::now();
        bubble_sort(&mut arr4);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert_eq!(arr4, vec![42]);

        // Test case 5: Array with duplicate elements
        let mut arr5 = vec![7, 5, 9, 5, 3, 7];
        let start5 = Instant::now();
        bubble_sort(&mut arr5);
        let elapsed5 = start5.elapsed();
        println!("Array with duplicate elements: {:?}", elapsed5);
        assert_eq!(arr5, vec![3, 5, 5, 7, 7, 9]);
    }
}
