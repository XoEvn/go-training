package singleton

import (
	"testing"

	singleton1"github.com/mohuishou/go-design-pattern/01_singleton"

	"github.com/stretchr/testify/assert"
)

func  TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton1.GetLazyInstance(),singleton1.GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton1.GetLazyInstance() != singleton1.GetLazyInstance(){
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton1.GetInstance(), singleton1.GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton1.GetInstance() != singleton1.GetInstance(){
				b.Errorf("test fail")
		}
	}
})
}



