#include "head/listnode.h"

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode* curA = headA;
        ListNode* curB = headB;
        int lenA = 0, lenB = 0;
        while (curA != nullptr) { // 求链表A的长度
            lenA++;
            curA = curA->next;
        }
        while (curB != nullptr) { // 求链表B的长度
            lenB++;
            curB = curB->next;
        }
        curA = headA;
        curB = headB;
        ListNode* fast;
        ListNode* slow;
        int gap;

        if (lenA > lenB) {
            gap = lenA-lenB;
            slow = curB;
            fast = curA;
        } else {
            gap = lenB-lenA;
            slow = curA;
            fast = curB;
        }

        for (int i = 0; i < gap; i++) {
            fast = fast->next;
        }

        while(fast != slow) {
            fast = fast->next;
            slow = slow->next;
        }
        return fast;
    }
};