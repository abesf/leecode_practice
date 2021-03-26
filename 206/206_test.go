package test206

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseList(t *testing.T) {
	a:=&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,nil}}}}}
	t.Log(reverseList(a))
}

//1->2->3->4->5->nil
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		node := head.Next   //2->3->4->5         3->4->5
		head.Next = newHead //1->nil             2->1->nil
		newHead = head      //newHead=1->nil     newHead=2->1->nil
		head = node         //head=2->3->4->5    head=3->4->5
	}
	return newHead
}
   
