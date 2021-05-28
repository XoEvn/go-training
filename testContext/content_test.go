package testContext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)

		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n:=range gen(ctx){
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}
