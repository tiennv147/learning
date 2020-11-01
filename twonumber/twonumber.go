package main

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 9,
	}

	l3 := addTwoNumbers(l1, l2)

	for {
		if l3 == nil {
			break
		}

		println("v", l3.Val)

		l3 = l3.Next
	}
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}

	carry := 0
	current := result
	last := result

	for {
		if current == nil {
			if carry > 0 {
				last.Next = &ListNode{
					Val:  carry,
					Next: nil,
				}
			}
			break
		}
		sum := l1.Val + l2.Val + carry

		current.Val = sum % 10
		carry = sum / 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			current.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		if l1 == nil {
			l1 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}
		if l2 == nil {
			l2 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		last = current
		current = current.Next
	}

	return result
}
