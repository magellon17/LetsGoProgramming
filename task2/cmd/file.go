package main

func findKthLargest(nums []int, k int) int {
	yanums := make([]int, len(nums))
	copy(yanums, nums)

	return quickSelect(yanums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := partition(nums, left, right)
	rank := pivotIndex - left + 1

	if k == rank {
		return nums[pivotIndex]
	} else if k < rank {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k-rank)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] >= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}
