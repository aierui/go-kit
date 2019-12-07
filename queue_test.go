package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.Empty(), "new queue is empty")

	s, e := 0, 100

	for i := s; i < e; i++ {
		q.Push(i)
		ast.Equal(i-s+1, q.Len(), "q push element")
	}

	for i := s; i < e; i++ {
		ast.Equal(i, q.Pop(), "q pop element")
	}

	ast.True(q.Empty(), "this queue is empty")
}
