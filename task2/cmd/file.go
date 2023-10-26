package main

func findKthLargest(nums []int, k int) int {
	kthLargest := make([]int, k)

	for i := 0; i < len(nums); i++ {
		if nums[i] > kthLargest[k-1] {
			for j := k - 1; j >= 1; j-- {
				kthLargest[j] = kthLargest[j-1]
			}
			kthLargest[0] = nums[i]
		}
	}

	return kthLargest[k-1]
}
