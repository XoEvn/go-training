package Array

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	var nums = []int{4, 4, 4, 5, 0, 0, 1, 3, 4}
	var target = 4
	result := twoSum(nums, target)
	for num, ok := range result {
		fmt.Printf("the index is  %x and the value  is %x\n", num, ok)
	}
}
// 两数之和
// 参数1: 目标查找的int数组
// 参数2: 目标和值
// 说明: 当存在目标之和的两个值,就将他们的下标查出来返回一个int数组,如果没有就直接返回一个nil
func twoSum(nums []int, target int) []int {
	// 初始化 map
	m := make(map[int]int)
	// 循环数组
	for i := 0; i < len(nums); i++ {
		// 获取减去的值
		another := target - nums[i]
		// 如果在map中有值,就获取值的下标,返回当前被减数的下标
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		// 将当前索引的值作为键,下标作为值 放入map中
		m[nums[i]] = i
	}
	return nil
}
