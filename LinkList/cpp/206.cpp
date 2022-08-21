#include "head/listnode.h"

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution
{
public:
  ListNode *reverseList(ListNode *head)
  {
    ListNode *temp;
    ListNode *cur = head;
    ListNode *pre = nullptr;

    while (cur != nullptr)
    {
      temp = cur->next;
      cur->next = pre;
      pre = cur;
      cur = temp;
    }
    return pre;
  }
};

// 递归
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class SolutionTwo {
public:
    ListNode* reverseList(ListNode* head) {
        return reverse(nullptr, head);
    }

    ListNode* reverse(ListNode* pre, ListNode* cur) {
      if (cur == nullptr) return pre;
      ListNode* tmp = cur->next;
      cur->next = pre;
      return reverse(cur, tmp);
    }
};