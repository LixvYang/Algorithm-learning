class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = nullptr;
        random = nullptr;
    }
};


class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (!head) return head;

        unordered_map<Node*, Node*> map;
        Node* cur = head;

        while(cur) {
            map[cur] = new Node(cur->val);
            cur = cur->next;
        }
        cur = head;
        while(cur) {
            map[cur]->next = map[cur->next];
            map[cur]->random = map[cur->random];
            cur = cur->next;
        }
        return map[head];
    }
};