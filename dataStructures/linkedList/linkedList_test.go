package linkedList

import "testing"

func TestLinkedList(t *testing.T) {
	//Test createNode function
	head := createNode('A')
	if head.value != 'A' {
		t.Errorf("error in createNode, want : 'A', got : %c", head.value)
	}

	//Test appendNode function
	appendedNode, _ := appendNode(head, 'B')
	if appendedNode.value != 'B' {
		t.Errorf("error in appended node value, want 'B', got : %c", appendedNode.value)
	}
	if appendedNode.prevVal != head {
		t.Errorf("error in appended node previous value, want %c, got : %c", head.value, appendedNode.prevVal.value)
	}

	newHead, _ := prependNode(head, 'Z')
	if newHead == nil || newHead.value != 'Z' {
		t.Errorf("error with the new head value, want : 'Z', got : %c", newHead.value)
	}
	if newHead.nextVal != head {
		t.Errorf("error with the nex node of the new head. Want : %c, got : %c", head.value, newHead.nextVal.value)
	}
	//TODO implement next tests
}
