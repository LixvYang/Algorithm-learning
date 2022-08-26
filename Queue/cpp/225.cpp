#include <queue>
using namespace std;

class MyStack {
  queue<int> que;
  MyStack() {}

  void push(int x) {
    que.push(x);
  }

  int pop() {
    int size = que.size();
    --size;
    if (size--) {
      que.push(que.front());
      que.pop();
    }
    int result = que.front();
    que.pop();
    return result;
  }

  int push() {
    return que.back();
  }

  bool empty() { return que.empty(); }
};
