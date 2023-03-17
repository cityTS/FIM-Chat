package tools

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	q := New()
	for i := 1; i <= 257; i++ {
		q.Push(i)
		fmt.Println("len: ", q.Size(), ",Cap: ", q.cap)
	}
}
func TestReduceCap(t *testing.T) {
	q := New()
	for i := 1; i <= 257; i++ {
		q.Push(i)
	}
	for i := 1; i <= 257; i++ {
		q.Pop()
		fmt.Println("len: ", q.Size(), ",Cap: ", q.cap)
	}
}

func TestQueue_Clear(t *testing.T) {
	q := New()
	for i := 1; i <= 257; i++ {
		q.Push(i)
	}
	fmt.Println(q)
	q.Clear()
	fmt.Println(q)
}

func TestQueue_Empty(t *testing.T) {
	q := New()
	fmt.Println(q.Empty())
	q.Push(11)
	fmt.Println(q.Empty())
}

func TestQueue_Front(t *testing.T) {
	q := New()
	fmt.Println(q.Front())
	q.Push(11)
	q.Push(22)
	fmt.Println(q.Front())
}

func TestQueue_Pop(t *testing.T) {
	q := New()
	q.Push(111)
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
}
