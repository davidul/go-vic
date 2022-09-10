package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Sample struct {
	n int
	s string
}

func TestLinkedList_Add(t *testing.T) {
	list := LinkedList{}
	size := list.Size()
	if size != 0 {
		t.Fatal("Wrong number of elements")
	}
	list.Add(1)
	size = list.Size()
	if size != 1 {
		t.Fatal("Wrong number of elements")
	}
	list.Add(2)
	size = list.Size()
	if size != 2 {
		t.Fatal("Wrong number of elements")
	}

}

func TestLinkedList_AddFirst(t *testing.T) {
	list := LinkedList{}
	list.AddFirst(1)
	size := list.Size()
	if size != 1 {
		t.Fatal("Wrong number of elements")
	}

	list.AddFirst(2)
	peek := list.Peek()
	if peek != 2 {
		t.Fatal("Failed")
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	list := LinkedList{}
	list.AddLast(1)
	size := list.Size()
	if size != 1 {
		t.Fatal("Wrong number of elements")
	}
	peek := list.Peek()
	if peek != 1 {
		t.Fatal("Peek failed")
	}

	list.AddLast(2)

	peek = list.Peek()
	if peek != 1 {
		t.Fatal("Peek failed")
	}

	size = list.Size()
	if size != 2 {
		t.Fatal("Wrong number of elements")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	intList := LinkedList{}
	structList := LinkedList{}

	intList.Add(1)

	structList.Add(Sample{
		n: 0,
		s: "test2",
	})
	structList.Add(Sample{
		n: 1,
		s: "test",
	})

	assert.True(t, intList.Contains(1))
	assert.False(t, intList.Contains(2))
	assert.True(t, structList.Contains(Sample{
		n: 1,
		s: "test",
	}))
}

func TestLinkedList_RemoveNode(t *testing.T) {
	structList := LinkedList{}
	t1 := structList.Add(Sample{
		n: 0,
		s: "test0",
	})
	t2 := structList.Add(Sample{
		n: 1,
		s: "test1",
	})
	t3 := structList.Add(Sample{
		n: 2,
		s: "test2",
	})
	t4 := structList.Add(Sample{
		n: 3,
		s: "test3",
	})

	structList.RemoveNode(t3)
	head := structList.Head()
	assert.Equal(t, head, t1)
	assert.Equal(t, head.next, t2)
	assert.Equal(t, head.next.next, t4)
	assert.Nil(t, head.next.next.next)

}
