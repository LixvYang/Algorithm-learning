#include "./head/listnode.h"

class Solution
{
public:
  ListNode *removeElements(ListNode *head, int val)
  {
    ListNode *dummyHead = new ListNode(0);
    dummyHead->next = head;
    ListNode *cur = dummyHead;
    while (cur->next != nullptr)
    {
      if (cur->next->val == val)
      {
        ListNode *tmp = cur->next;
        cur->next = cur->next->next;
        delete tmp;
      }
      else
      {
        cur = cur->next;
      }
    }
    head = dummyHead->next;
    delete dummyHead;
    return head;
  }
};