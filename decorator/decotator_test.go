package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColorSquare(t *testing.T) {
	sq :=Square{}
	csq :=NewColorSquare(sq,"red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red",got)
}
