package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		4,
		&ListNode{
			2,
			&ListNode{
				1,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}

	printList(list)
	//printList(reverseListRec(list))
	printList(reverseList(list))
}

func printList(head *ListNode) {
	d := ""
	p := head
	for p != nil {
		print(d, p.Val)
		d = "->"
		p = p.Next
	}
	println()
}

func reverseListRec(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	rList, _ := construct(head, nil)

	return rList
}

func construct(head, tail *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	if head.Next != nil {
		s, t := construct(head.Next, tail)
		t.Next = head
		head.Next = nil

		return s, head
	}

	return head, head
}

func reverseList(head *ListNode) *ListNode {
	var rList *ListNode

	l := head
	for l != nil {
		ln := l
		l = l.Next

		ln.Next = rList
		rList = ln
	}

	return rList
}
