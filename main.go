package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	maxSize := m + n                // Max size of merged array
	tempArr := make([]int, maxSize) // Create a new array to hold result
	i := 0                          // Set i to 0
	j := 0                          // Set j to 0
	index := 0                      // Set index to 0

	// Range through both array when both not reach the end of array
	for (i < m) && (j < n) {
		if nums1[i] < nums2[j] {
			tempArr[index] = nums1[i]
			i++
		} else {
			tempArr[index] = nums2[j]
			j++
		}
		index++
	}

	// Range through nums1 array if it has yet reach the end of array
	for i < m {
		tempArr[index] = nums1[i]
		i++
		index++
	}

	// Range through nums2 array if it has yet reach the end of array
	for j < n {
		tempArr[index] = nums2[j]
		j++
		index++
	}

	// Transfer result from tempArr to nums1 array
	for index = 0; index < m+n; index++ {
		nums1[index] = tempArr[index]
	}
}
