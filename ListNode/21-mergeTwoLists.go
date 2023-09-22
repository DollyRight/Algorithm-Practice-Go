package main

type LinkNode struct {
	Value    int64
	NextNode *LinkNode
}

func mergeTwoLists(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	//虚拟头节点，如果不用虚拟头节点，则还需额外考虑指针为空的情况，
	dummy := &LinkNode{-1, nil}
	p := dummy
	p1 := l1
	p2 := l2

	// 如果两个节点都存在，则通过比较把小的节点放进结果链表，把这个节点变为它的下一节点
	// 如果有节点不存在，则存放另一个存在的节点即可，这个节点后续的节点，就跟在这个节点后，一起进入结果链表了
	for p1 != nil && p2 != nil {
		// 比较p1和p2的两个指针
		// 将值较小的节点接到p指针
		if p1.Value > p2.Value {
			p.NextNode = p2
			p2 = p2.NextNode
		} else {
			p.NextNode = p1
			p1 = p1.NextNode
		}

		p = p.NextNode
	}

	if p1 != nil {
		p.NextNode = p1
	}

	if p2 != nil {
		p.NextNode = p2
	}

	return dummy.NextNode
}
