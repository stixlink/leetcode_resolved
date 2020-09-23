package _019__remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenList := 0
	index := 0
	head, _, _ = iterate(head, n, &lenList, &index)
	return head
}

func iterate(l *ListNode, n int, lenList, index *int) (*ListNode, *ListNode, bool) {
	*index++
	if l.Next == nil {
		*lenList = *index
		if *lenList == 1 && n == 1 {
			return nil, nil, false
		}
		return l, nil, true
	}
	nextCurrentNode, nextNextNode, isBk := iterate(l.Next, n, lenList, index)
	if isBk {
		*index--
		if *lenList-*index == n {
			l.Next = nextNextNode
		}
		if *lenList == n && *index == 1 {
			return nextCurrentNode, nextNextNode, false
		}
		return l, l.Next, true
	}

	return l, nextCurrentNode, true
}
