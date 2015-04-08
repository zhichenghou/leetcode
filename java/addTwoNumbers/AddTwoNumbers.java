/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class AddTwoNumbers {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int carry = 0;
        ListNode sum = new ListNode(-1);
        ListNode current = sum;
        
        while (l1 != null || l2 != null){
            int v1 = l1 == null ? 0 : l1.val;
            int v2 = l2 == null ? 0 : l2.val;
            int val = (v1 + v2 + carry) % 10 ;
           
            current.next = new ListNode(val);
            current = current.next;
        
            carry = (v1 + v2 + carry) / 10;
            l1 = l1 == null ?  null : l1.next;
            l2 = l2 == null ?  null : l2.next;
        }
        
        if (carry > 0){
            current.next = new ListNode(carry);
        }
        
        return sum.next;
        
    }
}