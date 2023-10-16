package linkedList

import "log"

type Node struct {
	value   rune
	nextVal *Node
	prevVal *Node
}

// createNode: create a new node with a value
func createNode(value rune) *Node {
	return &Node{
		value: value,
	}
}

// appendNode: add a node to the end of the list and return the head
func appendNode(head *Node, value rune) *Node {
	newNode := createNode(value)
	if head == nil {
		return newNode
	}

	currentNode := head
	for currentNode.nextVal != nil {
		currentNode = currentNode.nextVal
	}

	currentNode.nextVal = newNode
	newNode.prevVal = currentNode

	return head
}

func prependNode(head, newHead *Node) *Node {
	if head == nil {
		log.Fatal("There is no head.")
	}

	head.prevVal = newHead
	newHead.nextVal = head

	return newHead
}

// insertNode: insert a new Node
func insertNode(insertAfterNode, nodeToAdd *Node) *Node {
	if insertAfterNode == nil {
		log.Fatal("There is no Node to insert after.")
	}

	insertBeforeNode := insertAfterNode.nextVal

	nodeToAdd.prevVal = insertAfterNode
	nodeToAdd.nextVal = insertBeforeNode
	insertBeforeNode.prevVal = nodeToAdd
	insertAfterNode.nextVal = nodeToAdd

	return nodeToAdd
}

// deleteNode: delete a node
func deleteNode(nodeToDelete *Node) {
	prevNode := nodeToDelete.prevVal
	nextNode := nodeToDelete.nextVal

	prevNode.nextVal = nextNode
	nextNode.prevVal = prevNode

	nodeToDelete.prevVal = nil
	nodeToDelete.nextVal = nil
}
