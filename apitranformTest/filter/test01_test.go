package filter

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T){
	//f64, err := strconv.ParseFloat(value, 64)

	s := []int {7,8,9}
	p :=&s
	r :=*p
	r[0]=11
	fmt.Println(s[0])
}
