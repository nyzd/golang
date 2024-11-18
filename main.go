package main

import "fmt"

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

func main() {
	node := newNode(1, newNode(2, newNode(3, nil)))
	node.printAll()
	fmt.Println("-----")
	node.appendToEnd(4)
	node.printAll()

	fmt.Println("Doubly Linked List")

	dNode := &DoublyLinkedNode{
		value: 1,
		prev:  nil,
		next:  nil,
	}

	dNode.appendToEnd(2)

	dNode.printAll()
}
