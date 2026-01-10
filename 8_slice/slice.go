package main

import "fmt"

func main() {

	//var a []int
	//fmt.Println(len(a))

	//var a = make([]int, 0, 5)
	//fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))
	//a = append(a, 1)
	//a = append(a, 1)
	//a = append(a, 1)
	//a = append(a, 1)
	//a = append(a, 1)
	//a = append(a, 1)
	//fmt.Println(a)
	//fmt.Println(cap(a))

	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(nums)
	var nums = make([]int, 0, 5)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println(cap(nums))
	fmt.Println(nums[2:4])
}
