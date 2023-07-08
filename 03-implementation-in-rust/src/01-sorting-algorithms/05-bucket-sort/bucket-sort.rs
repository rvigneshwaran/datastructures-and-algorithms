fn bucket_sort(arr: &mut [f64]) {
    let n = arr.len(); // Get the length of the array
    if n <= 1 {
        return; // No need to sort if the array has 0 or 1 element
    }

    // Create buckets
    let mut buckets: Vec<Vec<f64>> = vec![vec![]; n];

    // Assign elements to buckets
    for &num in arr.iter() {
        let index = (n as f64 * num) as usize; // Calculate the index of the bucket
        buckets[index].push(num); // Assign the element to the corresponding bucket
    }

    // Sort individual buckets
    for bucket in buckets.iter_mut() {
        if bucket.len() > 1 {
            bucket.sort_unstable_by(|a, b| a.partial_cmp(b).unwrap()); // Sort each bucket using an unstable comparison
        }
    }

    // Concatenate sorted buckets into the original array
    let mut i = 0;
    for bucket in buckets.iter() {
        for &num in bucket.iter() {
            arr[i] = num; // Copy the sorted elements back into the original array
            i += 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::time::Instant;

    #[test]
    fn test_bucket_sort() {
        // Test case 1: Best-case scenario (Array with elements in ascending order)
        let mut arr1 = vec![0.1, 0.2, 0.3, 0.4, 0.5];
        let start1 = Instant::now();
        bucket_sort(&mut arr1);
        let elapsed1 = start1.elapsed();
        println!("Best-case scenario: {:?}", elapsed1);
        assert_eq!(arr1, vec![0.1, 0.2, 0.3, 0.4, 0.5]);

        // Test case 2: Worst-case scenario (Array with elements in descending order)
        let mut arr2 = vec![0.5, 0.4, 0.3, 0.2, 0.1];
        let start2 = Instant::now();
        bucket_sort(&mut arr2);
        let elapsed2 = start2.elapsed();
        println!("Worst-case scenario: {:?}", elapsed2);
        assert_eq!(arr2, vec![0.1, 0.2, 0.3, 0.4, 0.5]);

        // Test case 3: Empty array
        let mut arr3: Vec<f64> = vec![];
        let start3 = Instant::now();
        bucket_sort(&mut arr3);
        let elapsed3 = start3.elapsed();
        println!("Empty array: {:?}", elapsed3);
        assert_eq!(arr3, vec![]);

        // Test case 4: Array with a single element
        let mut arr4 = vec![42.0];
        let start4 = Instant::now();
        bucket_sort(&mut arr4);
        let elapsed4 = start4.elapsed();
        println!("Array with a single element: {:?}", elapsed4);
        assert_eq!(arr4, vec![42.0]);

        // Test case 5: Array with duplicate elements
        let mut arr5 = vec![0.7, 0.5, 0.9, 0.5, 0.3, 0.7];
        let start5 = Instant::now();
        bucket_sort(&mut arr5);
        let elapsed5 = start5.elapsed();
        println!("Array with duplicate elements: {:?}", elapsed5);
        assert_eq!(arr5, vec![0.3, 0.5, 0.5, 0.7, 0.7]);
    }
}
