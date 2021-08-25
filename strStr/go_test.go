/**
 * @Author: evnxo
 * @Description:
 * @File:  go_test
 * @Version: 1.0.0
 * @Date: 2021/8/22 2:04 下午
 */

package strStr

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	index := strings.Index("hello", "ll")
	fmt.Println(index)
}
