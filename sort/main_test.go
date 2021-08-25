/**
 * @Author: evnxo
 * @Description:
 * @File:  main_test.go
 * @Version: 1.0.0
 * @Date: 2021/8/22 2:58 下午
 */

package sort

import (
	"fmt"
	"testing"
)

func TestSortArray3(t *testing.T) {
	nums := sortArray3([]int{8, 5, 2, 1, 0})
	fmt.Println(nums)
}

func TestSortArray2(t *testing.T) {
	nums := sortArray2([]int{9, 9, 8, 2})
	fmt.Println(nums)
}

func TestSortArray(t *testing.T) {
	nums := sortArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	fmt.Println(nums)
}
