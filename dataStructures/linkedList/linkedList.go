package linkedList

import "errors"

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

// appendNode: add a node to the end of the list and return the added Node
func appendNode(head *Node, value rune) (*Node, error) {
	newNode := createNode(value)
	if head == nil {
		return nil, errors.New("you have to give a head or create one with createNode function")
	}

	currentNode := head
	for currentNode.nextVal != nil {
		currentNode = currentNode.nextVal
	}

	currentNode.nextVal = newNode
	newNode.prevVal = currentNode

	return newNode, nil
}

// prependNode: add a new head to the linked list and return this head
func prependNode(head *Node, newValue rune) (*Node, error) {
	if head == nil {
		return nil, errors.New("there is no head")
	}

	newHead := createNode(newValue)
	head.prevVal = newHead
	newHead.nextVal = head

	return newHead, nil
}

// insertNode: insert a new Node and returning the added node
func insertNode(insertAfterNode, nodeToAdd *Node) (*Node, error) {
	if insertAfterNode == nil {
		return nil, errors.New("there is no Node to insert after")
	}

	insertBeforeNode := insertAfterNode.nextVal

	nodeToAdd.prevVal = insertAfterNode
	nodeToAdd.nextVal = insertBeforeNode
	insertBeforeNode.prevVal = nodeToAdd
	insertAfterNode.nextVal = nodeToAdd

	return nodeToAdd, nil
}

// deleteNode: delete a node
func deleteNode(nodeToDelete *Node) {
	prevNode := nodeToDelete.prevVal
	nextNode := nodeToDelete.nextVal

	if prevNode == nil {
		nextNode.prevVal = nil
		nodeToDelete.nextVal = nil
	}

	prevNode.nextVal = nextNode
	nextNode.prevVal = prevNode

	nodeToDelete.prevVal = nil
	nodeToDelete.nextVal = nil
}
