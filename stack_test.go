package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {

	ast := assert.New(t)

	s := NewStack()

	ast.True(s.Empty(), "new stack is empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i-start+1, s.Len(), "stack length")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "pop element")
	}

	ast.True(s.Empty(), "this stack is empty")

}
