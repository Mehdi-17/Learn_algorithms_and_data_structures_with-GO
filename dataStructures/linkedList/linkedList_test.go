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
	appendedNode, _ := appendNode(head, 'B')
	if appendedNode.value != 'B' {
		t.Errorf("error in appended node value, want 'B', got : %c", appendedNode.value)
	}
	if appendedNode.prevVal != head {
		t.Errorf("error in appended node previous value, want %c, got : %c", head.value, appendedNode.prevVal.value)
	}

	thirdNode, _ := appendNode(head, 'C')
	if thirdNode.value != 'C' {
		t.Errorf("error in appended node value, want 'C', got : %c", thirdNode.value)
	}
	if thirdNode.prevVal != appendedNode {
		t.Errorf("error in appended node previous value, want %c, got : %c", appendedNode.value, thirdNode.prevVal.value)
	}
}

func TestPrependNode(t *testing.T) {
	head := createNode('A')
	newHead, _ := prependNode(head, 'Z')
	if newHead == nil || newHead.value != 'Z' {
		t.Errorf("error with the new head value, want : 'Z', got : %c", newHead.value)
	}
	if newHead.nextVal != head {
		t.Errorf("error with the nex node of the new head. Want : %c, got : %c", head.value, newHead.nextVal.value)
	}
}

func TestInsertNode(t *testing.T) {
	head := createNode('A')
	secondNode, _ := appendNode(head, 'B')
	thirdNode, _ := appendNode(head, 'C')

	insertedNode, _ := insertNode(secondNode, 'I')
	if insertedNode.prevVal != secondNode || insertedNode.nextVal != thirdNode {
		t.Errorf("Error during the insert, next and/or prev val are wrong.")
	}
}

func TestDeleteNode(t *testing.T) {
	head := createNode('A')
	secondNode, _ := appendNode(head, 'B')
	thirdNode, _ := appendNode(head, 'C')

	deleteNode(thirdNode)
	if secondNode.nextVal != nil {
		t.Errorf("Error during delete, the node is still in the linked list.")
	}
}

//TODO: add other specific test cases for insert and delete function :
// - delete and insert in the middle
// -delete and insert at the beginning

//TODO: Add test cases with failure + get err rather than _ on existing tests
