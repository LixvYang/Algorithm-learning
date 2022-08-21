class MyLinkedList
{
  struct LinkedNode
  {
    int val;
    LinkedNode *next;
    LinkedNode(int val) : val(val), next(nullptr) {}
  };

public:
  MyLinkedList()
  {
    size_ = 0;
    dummyHead_ = new LinkedNode(0);
  }

  int get(int index)
  {
    if (index < -1 || index > (size_ - 1))
    {
      return -1;
    }

    LinkedNode *cur = dummyHead_->next;
    while (index > 0)
    {
      index--;
      cur = cur->next;
    }
    return cur->val;
  }

  void addAtHead(int val)
  {
    LinkedNode *newNode = new LinkedNode(val);
    newNode->next = dummyHead_->next;
    dummyHead_->next = newNode;
    ++size_;
  }

  void addAtTail(int val)
  {
    LinkedNode *newNode = new LinkedNode(val);
    LinkedNode *cur = dummyHead_;
    while (cur->next != nullptr)
    {
      cur = cur->next;
    }
    cur->next = newNode;
    ++size_;
  }

  void addAtIndex(int index, int val)
  {
    if (index > size_ || index < 0)
    {
      return;
    }
    LinkedNode *newNode = new LinkedNode(val);
    LinkedNode *cur = dummyHead_;
    while (index--)
    {
      cur = cur->next;
    }
    newNode->next = cur->next;
    cur->next = newNode;
    ++size_;
  }

  void deleteAtIndex(int index)
  {
    if (index >= size_ || index < 0)
    {
      return;
    }
    LinkedNode *cur = dummyHead_;
    while (index--)
    {
      cur = cur->next;
    }
    LinkedNode *tmp = cur->next;
    cur->next = cur->next->next;
    delete tmp;
    --size_;
  }

private:
  int size_;
  LinkedNode *dummyHead_;
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList* obj = new MyLinkedList();
 * int param_1 = obj->get(index);
 * obj->addAtHead(val);
 * obj->addAtTail(val);
 * obj->addAtIndex(index,val);
 * obj->deleteAtIndex(index);
 */