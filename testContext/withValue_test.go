package testContext

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T)  {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v:= ctx.Value(k);v!=nil{
			fmt.Println("found value:",v)
			return
		}
		fmt.Println("key not found:",k)
	}
	k :=favContextKey("language")
	ctx :=context.WithValue(context.Background(),k,"GO")
	f(ctx, k)
	f(ctx, favContextKey("color"))

}
