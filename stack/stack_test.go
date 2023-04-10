package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[string]()
	assert.True(t, true, stack != nil)
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("World")
	stack.Push("Hello")
	pop := stack.Pop()
	result := fmt.Sprintf("%s %s", pop, stack.Pop())
	assert.Equal(t, "Hello World", result)
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")

	assert.Equal(t, "C", stack.Pop())
	assert.Equal(t, "B", stack.Pop())
	assert.Equal(t, "A", stack.Pop())
	assert.Equal(t, "", stack.Pop())
}

func BenchmarkStack_Push(b *testing.B) {
	stack := NewStack[string]()
	for n := 0; n < b.N; n++ {
		stack.Push("Hello")
	}
}
