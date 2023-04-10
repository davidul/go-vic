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

func TestLinkedList_AddNew(t *testing.T) {
	list := NewLinkedList("A")
	n1 := list.Add("B")
	n2 := list.Add("C")
	assert.Equal(t, list.count, 3)
	assert.Equal(t, n1.data, "B")
	assert.Equal(t, n2.data, "C")
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

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList[string]("A")
	list.Add("B")
	list.Add("C")
	r1 := list.Remove()
	r2 := list.Remove()
	r3 := list.Remove()
	r4 := list.Remove()
	r5 := list.Remove()

	assert.Equal(t, r1, "A")
	assert.Equal(t, r2, "B")
	assert.Equal(t, r3, "C")
	assert.Equal(t, r4, "")
	assert.Equal(t, r5, "")
}

func TestLinkedList_RemoveLast(t *testing.T) {
	list := NewLinkedList[string]("A")
	list.Add("B")
	list.Add("C")
	r1 := list.RemoveLast()
	r2 := list.RemoveLast()
	r3 := list.RemoveLast()
	r4 := list.RemoveLast()

	assert.Equal(t, "C", r1)
	assert.Equal(t, "B", r2)
	assert.Equal(t, "A", r3)
	assert.Equal(t, "", r4)
}

func TestLinkedList_Poll(t *testing.T) {
	list := NewLinkedList[string]("A")
	list.Add("B")
	list.Add("C")
	list.Add("D")

	p1 := list.Poll()
	p2 := list.Poll()
	p3 := list.Poll()
	p4 := list.Poll()

	assert.Equal(t, "A", p1)
	assert.Equal(t, "B", p2)
	assert.Equal(t, "C", p3)
	assert.Equal(t, "D", p4)
}

func TestLinkedList_Append(t *testing.T) {
	l1 := NewLinkedList[string]("A")
	l2 := NewLinkedList[string]("B")
	l1.Append(l2)
	p1 := l1.Poll()
	p2 := l1.Poll()
	assert.Equal(t, "A", p1)
	assert.Equal(t, "B", p2)

	l3 := NewLinkedList[string]("C")
	l3.Add("D")
	l3.Append(l2)
	l1.Append(l3)

	p3 := l3.Poll()
	p4 := l3.Poll()
	p5 := l3.Poll()

	assert.Equal(t, "C", p3)
	assert.Equal(t, "D", p4)
	assert.Equal(t, "", p5)

}

func BenchmarkLinkedList_Add(b *testing.B) {
	l := LinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.Add("hello")
	}
}

func BenchmarkLinkedList_AddFirst(b *testing.B) {
	l := LinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.AddFirst("hello")
	}
}

func BenchmarkLinkedList_AddLast(b *testing.B) {
	l := LinkedList[string]{}
	for n := 0; n < b.N; n++ {
		l.AddLast("hello")
	}
}

func BenchmarkLinkedList_Append(b *testing.B) {
	l := NewLinkedList[string]("Y")
	for n := 0; n < b.N; n++ {
		l1 := &LinkedList[string]{}
		l1.Add("X")
		l.Append(l1)
	}
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	l := NewLinkedList[int](-1)
	for i := 0; i < 1_000_000_0; i++ {
		l.Add(i)
	}

	b.ResetTimer()
	var r = 0
	for n := 0; n < b.N; n++ {
		r = l.Remove()
	}
	r = r
}
