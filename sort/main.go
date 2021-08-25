/**
 * @Author: evnxo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/8/22 2:58 下午
 */

package sort

func sortArray4(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + ((right - left) >> 1)
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	temparr := make([]int, right-left+1)
	temp1 := left
	temp2 := mid + 1
	var index int = 0
	for temp1 <= mid && temp2 <= right {
		if nums[temp1] <= nums[temp2] {
			index++
			temp1++
			temparr[index] = nums[temp1]
		} else {
			index++
			temp2++
			temparr[index] = nums[temp2]
		}
	}
	if temp1 <= mid {

	}
}

func sortArray3(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		var j int
		for j := i - 1; j >= 0; j-- {
			if temp < nums[j] {
				nums[j+1] = nums[j]
				continue
			}
			break
		}
		nums[j+1] = temp
	}
	return nums
}

// 简单选择
func sortArray2(nums []int) []int {
	len := len(nums)
	min := 0
	for i := 0; i < len; i++ {
		min = i
		for j := i + 1; j < len; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			swap(nums, i, min)
		}
	}
	return nums
}

// 冒泡排序
func sortArray(nums []int) []int {
	len := len(nums)
	flag := true
	for i := 0; i < len && flag; i++ {
		flag = false
		for j := 0; j < len-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				flag = true
			}
		}
	}
	return nums
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
