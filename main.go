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
	node *Node
}

func (s *LinkedListStack) push(value int) {
	s.node.appendToEnd(value)
}

type StackError struct {
	message string
}

func (s StackError) Error() string {
	return s.message
}

func (s *LinkedListStack) pop() (Node, error) {
	if s.node == nil {
		return *newNode(0, nil), StackError{"Stack is empty"}
	}
	return s.node.popLastNode(), nil
}

func main() {
	node := newNode(1, newNode(2, newNode(3, nil)))
	node.printAll()
	fmt.Println("-----")
	node.appendToEnd(4)
	node.printAll()

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
		node: newNode(1, nil),
	}

	stack.push(2)
	stack.push(3)

	val, err := stack.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(val.value)
}
