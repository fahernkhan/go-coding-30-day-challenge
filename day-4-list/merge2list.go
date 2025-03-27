package main

import "fmt"

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// rekursion
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// tabulation
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */
//  func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     if list1 == nil {
//         return list2
//     }
//     if list2 == nil {
//         return list1
//     }

//     dummy := &ListNode{} // Dummy head to simplify the logic
//     curr := dummy

//     for list1 != nil && list2 != nil {
//         if list1.Val < list2.Val {
//             curr.Next = list1
//             list1 = list1.Next
//         } else {
//             curr.Next = list2
//             list2 = list2.Next
//         }
//         curr = curr.Next
//     }

//     // Append the remaining nodes from the non-empty list
//     if list1 != nil {
//         curr.Next = list1
//     } else {
//         curr.Next = list2
//     }

//     return dummy.Next // Return the head of the merged list
// }

// Helper function to create a linked list from an array
func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	list1 := createLinkedList([]int{1, 2, 4})
	list2 := createLinkedList([]int{1, 3, 4})
	mergedList := mergeTwoLists(list1, list2)
	printLinkedList(mergedList) // Output: 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil

	list3 := createLinkedList([]int{})
	list4 := createLinkedList([]int{})
	mergedList2 := mergeTwoLists(list3, list4)
	printLinkedList(mergedList2) // Output: nil

	list5 := createLinkedList([]int{})
	list6 := createLinkedList([]int{0})
	mergedList3 := mergeTwoLists(list5, list6)
	printLinkedList(mergedList3) // Output: 0 -> nil
}
