package main

import (
	"fmt"
	"slices"
)

func main() {
	//dynamic array
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// nums[0] = 3
	// nums[1] = 5

	// nums = append(nums, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// //copy function

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[2:])

	//slice

	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}
	fmt.Println(slices.Equal(nums1, nums2))
}
