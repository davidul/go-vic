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
	list := DoublyLinkedList[int]{}
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
	list := DoublyLinkedList[int]{}
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
	list := DoublyLinkedList[int]{}
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
	intList := DoublyLinkedList[int]{}
	structList := DoublyLinkedList[Sample]{}

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
	structList := DoublyLinkedList[Sample]{}
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
	//t1  t2   t4   tail -> nil
	assert.Nil(t, head.next.next.next)

}

func TestLinkedList_Compare(t *testing.T) {
	structList := DoublyLinkedList[Sample]{}
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

	structList2 := DoublyLinkedList[Sample]{}
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
	l1 := DoublyLinkedList[string]{}
	l2 := DoublyLinkedList[string]{}
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

func TestLinkedList_Remove(t *testing.T) {
	list := NewDoublyLinkedList[string]("A")
	list.Add("X")
	list.Add("Y")
	r1 := list.Remove()
	r2 := list.Remove()
	r3 := list.Remove()
	r4 := list.Remove()

	assert.Equal(t, r1, "A")
	assert.Equal(t, r2, "X")
	assert.Equal(t, r3, "Y")
	assert.Equal(t, r4, "")

}

func BenchmarkLinkedList_Add(b *testing.B) {
	l := DoublyLinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.Add("hello")
	}
}

func BenchmarkLinkedList_AddFirst(b *testing.B) {
	l := DoublyLinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.AddFirst("hello")
	}
}

func BenchmarkLinkedList_AddLast(b *testing.B) {
	l := DoublyLinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.AddLast("hello")
	}
}

func BenchmarkLinkedList_Append(b *testing.B) {
	l := NewDoublyLinkedList[string]("Y")
	for n := 0; n < b.N; n++ {
		l1 := &DoublyLinkedList[string]{}
		l1.Add("X")
		l.Append(l1)
	}
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	l := NewDoublyLinkedList[int](-1)
	for i := 0; i < 1000_000; i++ {
		l.Add(i)
	}

	b.ResetTimer()
	var r = 0
	for n := 0; n < b.N; n++ {
		r = l.Remove()
	}
	r = r
}
