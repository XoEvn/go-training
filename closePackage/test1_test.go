package closePackage

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		c1 := createCounter(9)
		fmt.Println(c1())
		fmt.Println(c1())
	})
}
