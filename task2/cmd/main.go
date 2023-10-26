package main

func findKthLargest(arr []int, k int) int {
	kthLargest := make([]int, k)

	for i := 0; i < len(arr); i++ {
		if arr[i] > kthLargest[k-1] {
			for j := k - 1; j >= 1; j-- {
				kthLargest[j] = kthLargest[j-1]
			}
			kthLargest[0] = arr[i]
		}
	}

	return kthLargest[k-1]
}
