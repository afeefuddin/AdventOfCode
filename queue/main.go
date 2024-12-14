package queue

type node[T any] struct {
	val  T
	next *node[T]
}

type Queue[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func MakeQueue[T any]() *Queue[T] {
	qu := &Queue[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
	return qu
}

func (qu *Queue[T]) IsEmpty() bool {
	return qu.size == 0
}

func (qu *Queue[T]) Push(val T) {
	newNode := &node[T]{
		val:  val,
		next: nil,
	}

	if qu.IsEmpty() {
		qu.head = newNode
		qu.tail = newNode
	} else {
		qu.tail.next = newNode
		qu.tail = newNode
	}
	qu.size++
}

func (qu *Queue[T]) Pop() T {
	if qu.IsEmpty() {
		panic("Queue already Empty")
	}
	val := qu.head.val

	if qu.size == 1 {
		qu.head = nil
		qu.tail = nil
	} else {
		qu.head = qu.head.next
	}

	qu.size--

	return val
}
