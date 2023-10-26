package main

func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1 // Вернуть -1, если значение k находится за пределами допустимого диапазона
	}

	return quickSelect(nums, 0, len(nums)-1, k)
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
