/**
 * @Author: evnxo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/8/22 1:53 下午
 */

package strStr

import "strings"

func strStr(haystack, needle string) int {
	haylen := len(haystack)
	needlen := len(needle)
	if haylen < needlen {
		return -1
	}
	if needlen == 0 {
		return 0
	}

	if strings.Contains(haystack, needle) {

	}

	return 0

}
