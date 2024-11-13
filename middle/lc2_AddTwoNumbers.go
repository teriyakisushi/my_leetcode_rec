package main

import "fmt"

/*
@ Problem: Leetcode 2. Add Two Numbers
@ Solution: 模拟
@ Time Complexity: O(n)
@ Detail:
	两个链表相加，注意进位即可。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head1 := l1
    head2 := l2
    
    VirList := &ListNode{} // 定一个虚拟List
    curList := VirList //当前节点
    carry := 0 //处理进位

    for head1 !=nil || head2 !=nil{
        var x,y int
        if head1 !=nil{
            x = head1.Val
            head1=head1.Next  //到下一位
        }
        if head2 !=nil{
            y = head2.Val
            head2=head2.Next //到下一位
        }

        sum:= carry+x+y  //将当前的数字和加起来，包括上一次的进位
        carry = sum/10   //当前进位值
        curList.Next = &ListNode{Val: sum%10} //创建新节点保存当前为结果
        curList = curList.Next //到下一节点
        }

        if carry>0{
            curList.Next = &ListNode{Val: carry}
        }
        return VirList.Next
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}