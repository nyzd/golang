package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func newNode(value int, next *Node) *Node {
	return &Node{
		value,
		next,
	}
}

func (n *Node) print() {
	if n == nil {
		fmt.Println("Node Not found")
		return
	}
	fmt.Printf("node(%p -> %p): %d\n", n, n.next, n.value)
}

func (n *Node) printAll() {
	n.print()
	if n.next != nil {
		n.next.printAll()
	}
}

func (n *Node) linearSearch(value int) *Node {
	if n.value == value {
		return n
	} else if n.next != nil {
		return n.next.linearSearch(value)
	}

	return nil
}

func (n *Node) popLastNode() Node {
	if n.next == nil {
		curr := *n

		n = nil

		return curr
	}
	if n.next.next == nil {
		next := *n.next
		n.next = nil
		return next
	} else {
		return n.next.popLastNode()
	}
}

func (n *Node) reverse() *Node {
	curr := n
	var prev *Node = nil

	for curr != nil {
		next := curr.next
		curr.next = prev

		prev = curr
		curr = next
	}

	return prev
}

// Return's appended Node
// Linear Append O(n)
func (n *Node) appendToEnd(value int) *Node {
	if n.next == nil {
		n.next = newNode(value, nil)
		return n.next
	} else {
		return n.next.appendToEnd(value)
	}
}

type DoublyLinkedNode struct {
	value int
	next  *DoublyLinkedNode
	prev  *DoublyLinkedNode
}

func (n *DoublyLinkedNode) appendToEnd(value int) *DoublyLinkedNode {
	if n.next == nil {
		n.next = &DoublyLinkedNode{
			value: value,
			prev:  n,
			next:  nil,
		}
		return n.next
	} else {
		return n.next.appendToEnd(value)
	}
}

func (n *DoublyLinkedNode) print() {
	if n == nil {
		fmt.Println("Node Not found")
		return
	}
	fmt.Printf("node(%p <- %p -> %p): %d\n", n.prev, n, n.next, n.value)
}

func (n *DoublyLinkedNode) printAll() {
	n.print()
	if n.next != nil {
		n.next.printAll()
	}
}

type LinkedListStack struct {
	head *Node
}

func (s *LinkedListStack) push(value int) {
	new_node := newNode(value, s.head)
	s.head = new_node
}

type StackError struct {
	message string
}

func (s StackError) Error() string {
	return s.message
}

func (s *LinkedListStack) pop() (Node, error) {
	poped := *s.head

	s.head = s.head.next

	return poped, nil
}

type Queue struct {
	data    [10]int
	head    int
	tail    int
	last_op int8
}

func (q *Queue) enqueue(element int) {
	if q.tail == q.head && q.last_op == 0 {
		panic("Queue is full")
	}

	q.data[q.tail] = element
	if q.tail == 10-1 {
		q.tail = 0
	} else {
		q.tail += 1
	}

	q.last_op = 0
}

func (q *Queue) dequeue() int {
	if q.tail == q.head && q.last_op == 1 {
		panic("Queue is empty")
	}

	o := q.data[q.head]

	q.data[q.head] = 0
	q.head += 1

	q.last_op = 1

	return o
}

func main() {
	node := newNode(1, newNode(2, newNode(3, nil)))
	node.printAll()
	fmt.Println("-----")
	node.appendToEnd(4)
	node.printAll()

	fmt.Println("[* LL] Reversed.")
	node.reverse().printAll()

	fmt.Println("[*] Doubly Linked List")

	dNode := &DoublyLinkedNode{
		value: 1,
		prev:  nil,
		next:  nil,
	}

	dNode.appendToEnd(2)

	dNode.printAll()

	fmt.Println("[*] Stack")
	stack := LinkedListStack{
		head: newNode(1, nil),
	}

	stack.push(2)
	stack.push(3)

	stack.head.printAll()

	val, _ := stack.pop()
	val2, _ := stack.pop()
	val3, _ := stack.pop()

	fmt.Println(val.value)
	fmt.Println(val2.value)
	fmt.Println(val3.value)

	fmt.Println("[*] Queue")

	q := Queue{data: [10]int{}, head: 0, tail: 0, last_op: 1}

	q.enqueue(5)

	fmt.Println(q.data)
}
