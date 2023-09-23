package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}

	p1 := dummy1
	p2 := dummy2

	p := head
	for p != nil {
		fmt.Println(p.Val)
		if p.Val < x {
			//p1 下一节点变为p的首节点
			//p1 移到下一节点
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}

		//暂存p的下一节点
		temp := p.Next
		//断开p与p的下一节点
		p.Next = nil
		//p变成暂存的p的下一节点
		p = temp

	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
