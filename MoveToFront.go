package main

import "fmt"

type Element int

type node struct {
	data Element
	next *node
}

func Prev_Sequential_Search(head *node, target Element) *node {

	if head == nil {
		return nil
	}

	current := head

	for current.next != nil {
		if current.next.data == target {
			break
		}
		current = current.next
	}
	return current
}

func SLL_MoveToFront(head **node, data Element) {
	previousNode := Prev_Sequential_Search(*head, data)
	targetNode := previousNode.next
	if previousNode == nil {
		fmt.Println("There is no data you search")
		return
	}
	fmt.Println("search : ", previousNode.next.data)
	previousNode.next = targetNode.next
	targetNode.next = *head
	*head = targetNode
}

func PrintList(head *node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func InputNode(head **node, data Element) {
	newNode := &node{data: data, next: nil}
	current := *head
	if current == nil {
		*head = newNode
		return
	}
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func main() {
	var head *node = nil
	InputNode(&head, 71)
	InputNode(&head, 5)
	InputNode(&head, 13)
	InputNode(&head, 1)
	InputNode(&head, 2)
	InputNode(&head, 48)
	InputNode(&head, 222)
	InputNode(&head, 136)
	InputNode(&head, 3)
	InputNode(&head, 15)
	PrintList(head)
	fmt.Println("after search")
	SLL_MoveToFront(&head, 48)
	PrintList(head)
}
