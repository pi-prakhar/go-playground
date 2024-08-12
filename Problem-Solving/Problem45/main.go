package main

/**
PROBLEM:
leetcode -> Merge Sorted Array -> medium

TAGS:
array, two pointe
**/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		pointer3 := m + n - 1
		for pointer3 >= 0 {
			nums1[pointer3] = nums2[n-1]
			pointer3--
			n--
		}
		return
	}
	pointer1 := m - 1
	pointer2 := n - 1
	pointer3 := m + n - 1
	for pointer1 >= 0 && pointer2 >= 0 {
		if nums1[pointer1] > nums2[pointer2] {
			nums1[pointer3] = nums1[pointer1]
			pointer1--
		} else {
			nums1[pointer3] = nums2[pointer2]
			pointer2--
		}
		pointer3--
	}
	for pointer2 >= 0 {
		nums1[pointer2] = nums2[pointer2]
		pointer2--
	}

	return
}
