package main

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type List[T comparable] struct {
	head *Node[T]
}

// adds a new element to the end
func (l *List[T]) Add(item T) {
	if l.head == nil {
		l.head = &Node[T]{value: item}
		return
	}

	iterator := l.head
	for {
		if iterator.next == nil {
			iterator.next = &Node[T]{value: item}
			return
		}
		iterator = iterator.next
	}
}

// adds an element in the specified position
func (l *List[T]) Insert(item T, pos int) error {
	if pos < 0 {
		return fmt.Errorf("Must be a positive value. Use zero for head")
	}
	iterator := l.head
	if pos == 0 {
		l.head = &Node[T]{value: item, next: iterator}
		return nil
	}
	place := 0
	// We want to insert AT that location so stop one short
	for place < pos-1 && iterator != nil {
		place++
		iterator = iterator.next
	}
	if iterator == nil {
		return fmt.Errorf("That position is too far! The list is only %d items long\n", place)
	}
	oldNext := iterator.next
	iterator.next = &Node[T]{value: item, next: oldNext}
	return nil
}

// returns the position of the supplied value or -1 if not present
func (l *List[T]) Index(findMe T) int {
	iterator := l.head
	place := 0
	for iterator != nil {
		if iterator.value == findMe {
			return place
		}
		iterator = iterator.next
		place++
	}
	return -1
}

func (l *List[T]) Print() {
	var iterator *Node[T] = l.head
	index := 0
	for iterator != nil {
		fmt.Printf("%d - %v\n", index, iterator.value)
		iterator = iterator.next
		index++
	}
}
