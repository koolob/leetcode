package add_two_numbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	retCus := ret
	var parCus *ListNode = nil
	l1Cus := l1
	l2Cus := l2
	add1 := false
	for ; l1Cus != nil && l2Cus != nil; {
		retCus.Val = l1Cus.Val + l2Cus.Val
		if add1 {
			retCus.Val = retCus.Val + 1
			add1 = false
		}
		if retCus.Val >= 10 {
			add1 = true
			retCus.Val = retCus.Val - 10
		}
		parCus = retCus
		retCus.Next = &ListNode{}
		retCus = retCus.Next
		l1Cus = l1Cus.Next
		l2Cus = l2Cus.Next
	}
	for ; l1Cus != nil; {
		if add1 {
			retCus.Val = l1Cus.Val + 1
			add1 = false
		} else {
			retCus.Val = l1Cus.Val
		}
		if retCus.Val >= 10 {
			add1 = true
			retCus.Val = retCus.Val - 10
		}
		parCus = retCus
		retCus.Next = &ListNode{}
		retCus = retCus.Next
		l1Cus = l1Cus.Next
	}
	for ; l2Cus != nil; {
		if add1 {
			retCus.Val = l2Cus.Val + 1
			add1 = false
		}else {
			retCus.Val = l2Cus.Val
		}
		if retCus.Val >= 10 {
			add1 = true
			retCus.Val = retCus.Val - 10
		}
		parCus = retCus
		retCus.Next = &ListNode{}
		retCus = retCus.Next
		l2Cus = l2Cus.Next
	}
	if add1 {
		retCus.Val = 1
		retCus.Next = nil
	} else {
		parCus.Next = nil
	}

	return ret
}

func Test() {

	l1 := &ListNode{
		1, &ListNode{
			8, nil,
		},
	}
	l2 := &ListNode{
		0, nil,
	}
	t := addTwoNumbers(l1, l2)
	echo(t)
}

func echo(node *ListNode) {
	one := node
	for one != nil {
		fmt.Print(one.Val)
		one = one.Next
	}
	fmt.Println("\n")
}
