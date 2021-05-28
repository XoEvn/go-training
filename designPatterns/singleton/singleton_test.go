package singleton

import (
	"testing"

	singleton "github.com/mohuishou/go-design-pattern/01_singleton"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if singleton.GetInstance()!=singleton.GetInstance()}
		b.Error("test fail")
		}
	})
}
