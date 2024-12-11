package linkedlist

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func MakeLL[T any](arr []T) *Node[T] {
	if len(arr) == 0 {
		return nil
	}
	head := &Node[T]{
		Val:  arr[0],
		Next: nil,
	}

	ptr := head
	for i, val := range arr {
		if i == 0 {
			continue
		}
		temp := &Node[T]{
			Val:  val,
			Next: nil,
		}
		ptr.Next = temp
		ptr = ptr.Next
	}
	return head
}

func PrintLL[T any](head *Node[T]) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Print("end\n")
}

func InsertForward[T any](head *Node[T], value T) {
	temp := &Node[T]{
		Val:  value,
		Next: head.Next,
	}
	(*head).Next = temp
}

func Size[T any](head *Node[T]) int {
	size := 0
	for head != nil {
		head = head.Next
		size++
	}
	return size
}
