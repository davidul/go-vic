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
	list := LinkedList[int]{}
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
	list := LinkedList[int]{}
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
	list := LinkedList[int]{}
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
	intList := LinkedList[int]{}
	structList := LinkedList[Sample]{}

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
	structList := LinkedList[Sample]{}
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

func TestLinkedList_Compare(t *testing.T) {
	structList := LinkedList[Sample]{}
	structList.Add(Sample{
		n: 0,
		s: "test0",
	})
	structList.Add(Sample{
		n: 1,
		s: "test1",
	})
	structList.Add(Sample{
		n: 2,
		s: "test2",
	})
	structList.Add(Sample{
		n: 3,
		s: "test3",
	})

	structList2 := LinkedList[Sample]{}
	structList2.Add(Sample{
		n: 0,
		s: "test0",
	})
	structList2.Add(Sample{
		n: 1,
		s: "test1",
	})
	structList2.Add(Sample{
		n: 2,
		s: "test2",
	})
	structList2.Add(Sample{
		n: 3,
		s: "test3",
	})

	assert.True(t, structList.Compare(&structList2))
	assert.True(t, structList2.Compare(&structList))

	structList.Poll()
	assert.False(t, structList.Compare(&structList2))
	assert.False(t, structList2.Compare(&structList))

	structList2.Poll()
	assert.True(t, structList.Compare(&structList2))
}

func TestLinkedList_AddAll(t *testing.T) {
	l1 := LinkedList[string]{}
	l2 := LinkedList[string]{}
	l1.Add("Hello")
	l1.Add("World")
	l2.Add("How")
	l2.Add("Are")
	l2.Add("You")
	l1.AddAll(&l2)

	assert.Equal(t, l1.Size(), 5)
	assert.Equal(t, l1.Head().Data(), "Hello")
	assert.Equal(t, l1.Head().Next().Data(), "World")
	assert.Equal(t, l1.Head().Next().Next().Data(), "How")
	assert.Equal(t, l1.Head().Next().Next().Next().Data(), "Are")
	assert.Equal(t, l1.Head().Next().Next().Next().Next().Data(), "You")
	assert.Nil(t, l1.Head().Next().Next().Next().Next().Next())

}
