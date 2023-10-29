package linkedList

import "testing"

func TestCreateNode(t *testing.T) {
	//Test createNode function
	head := createNode('A')
	if head.value != 'A' {
		t.Errorf("error in createNode, want : 'A', got : %c", head.value)
	}
}

func TestAppendNode(t *testing.T) {
	head := createNode('A')
	//Test appendNode function
	appendedNode, errB := appendNode(head, 'B')

	if errB != nil {
		t.Errorf("appendNode(head, 'B') return an error. %v", errB)
	} else if appendedNode.value != 'B' {
		t.Errorf("error in appended node value, want 'B', got : %c", appendedNode.value)
	} else if appendedNode.prevVal != head {
		t.Errorf("error in appended node previous value, want %c, got : %c", head.value, appendedNode.prevVal.value)
	}

	thirdNode, errC := appendNode(head, 'C')

	if errC != nil {
		t.Errorf("appendNode(head, 'C') return an error. %v", errC)
	} else if thirdNode.value != 'C' {
		t.Errorf("error in appended node value, want 'C', got : %c", thirdNode.value)
	}
	if thirdNode.prevVal != appendedNode {
		t.Errorf("error in appended node previous value, want %c, got : %c", appendedNode.value, thirdNode.prevVal.value)
	}
}

func TestAppendNode_returnErr(t *testing.T) {
	appendedNode, err := appendNode(nil, 'B')
	if err == nil {
		t.Errorf("appendNode(nil, 'B') should return an error. %v", err)
	} else if appendedNode != nil {
		t.Errorf("Call appendNode with nil head should throw an error.")
	}
}

func TestPrependNode(t *testing.T) {
	head := createNode('A')
	newHead, err := prependNode(head, 'Z')

	if err != nil {
		t.Errorf("prependNode(head, 'Z'), return an error. %v", err)
	} else if newHead == nil || newHead.value != 'Z' {
		t.Errorf("error with the new head value, want : 'Z', got : %c", newHead.value)
	}
	if newHead.nextVal != head {
		t.Errorf("error with the nex node of the new head. Want : %c, got : %c", head.value, newHead.nextVal.value)
	}
}

func TestInsertNode(t *testing.T) {
	head := createNode('A')
	secondNode, errB := appendNode(head, 'B')
	thirdNode, errC := appendNode(head, 'C')

	if errB != nil || errC != nil {
		t.Errorf("One of the appendNode method return an error. %v | %v", errB, errC)
	}

	insertedNode, errInsert := insertNode(secondNode, 'I')

	if errInsert != nil {
		t.Errorf("insertNode(secondNode, 'I') return an error. %v", errInsert)
	} else if insertedNode.prevVal != secondNode || insertedNode.nextVal != thirdNode {
		t.Errorf("Error during the insert, next and/or prev val are wrong.")
	}
}

func TestDeleteNode(t *testing.T) {
	head := createNode('A')
	secondNode, errB := appendNode(head, 'B')
	thirdNode, errC := appendNode(head, 'C')

	if errB != nil || errC != nil {
		t.Errorf("One of the appendNode method return an error. %v | %v", errB, errC)
	}

	deleteNode(thirdNode)
	if secondNode.nextVal != nil || thirdNode.prevVal != nil || thirdNode.nextVal != nil {
		t.Errorf("Error during delete, the node is still in the linked list.")
	}
}

func TestDeleteNodeAtTheBeginning(t *testing.T) {
	head := createNode('A')
	secondNode, errB := appendNode(head, 'B')
	_, errC := appendNode(head, 'C')

	if errB != nil || errC != nil {
		t.Errorf("One of the appendNode method return an error. %v | %v", errB, errC)
	}

	deleteNode(head)
	if secondNode.prevVal != nil || head.prevVal != nil || head.nextVal != nil {
		t.Errorf("Error during delete, the node is still in the linked list.")
	}
}
