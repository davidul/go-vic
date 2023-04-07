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
