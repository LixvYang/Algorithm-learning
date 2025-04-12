/*
 * 设计链表
 * https://leetcode.cn/problems/design-linked-list/
 *
 * 题目描述：
 * 设计链表的实现。您可以选择使用单链表或双链表。
 * 单链表中的节点应该具有两个属性：val 和 next。
 * val 是当前节点的值，next 是指向下一个节点的指针/引用。
 * 如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。
 * 假设链表中的所有节点都是 0-index 的。
 *
 * 在链表类中实现这些功能：
 * - get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
 * - addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
 * - addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
 * - addAtIndex(index, val)：在链表中的第 index 个节点之前添加值为 val 的节点。
 *   如果 index 等于链表的长度，则该节点将附加到链表的末尾。
 *   如果 index 大于链表长度，则不会插入节点。
 * - deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 *
 * 示例：
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1, 2);  // 链表变为1-> 2-> 3
 * linkedList.get(1);            // 返回2
 * linkedList.deleteAtIndex(1);  // 现在链表是1-> 3
 * linkedList.get(1);            // 返回3
 *
 * 提示：
 * 所有val值都在 [1, 1000] 之内。
 * 操作次数将在  [1, 1000] 之内。
 * 请不要使用内置的 LinkedList 库。
 */

// 定义链表节点
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

/// 设计链表实现
struct MyLinkedList {
    head: Option<Box<ListNode>>,
    size: i32,
}

impl MyLinkedList {
    /// 初始化链表
    fn new() -> Self {
        MyLinkedList {
            head: None,
            size: 0,
        }
    }

    /// 获取链表中第 index 个节点的值
    fn get(&self, index: i32) -> i32 {
        if index < 0 || index >= self.size {
            return -1;
        }

        let mut curr = self.head.as_ref();
        for _ in 0..index {
            curr = curr.unwrap().next.as_ref();
        }
        curr.unwrap().val
    }

    /// 在链表头部添加节点
    fn add_at_head(&mut self, val: i32) {
        let mut new_node = Box::new(ListNode::new(val));
        new_node.next = self.head.take();
        self.head = Some(new_node);
        self.size += 1;
    }

    /// 在链表尾部添加节点
    fn add_at_tail(&mut self, val: i32) {
        let new_node = Box::new(ListNode::new(val));
        if self.head.is_none() {
            self.head = Some(new_node);
        } else {
            let mut curr = self.head.as_mut().unwrap();
            while curr.next.is_some() {
                curr = curr.next.as_mut().unwrap();
            }
            curr.next = Some(new_node);
        }
        self.size += 1;
    }

    /// 在指定位置添加节点
    fn add_at_index(&mut self, index: i32, val: i32) {
        if index > self.size {
            return;
        }

        if index <= 0 {
            self.add_at_head(val);
            return;
        }

        if index == self.size {
            self.add_at_tail(val);
            return;
        }

        let mut curr = self.head.as_mut().unwrap();
        for _ in 0..index - 1 {
            curr = curr.next.as_mut().unwrap();
        }

        let mut new_node = Box::new(ListNode::new(val));
        new_node.next = curr.next.take();
        curr.next = Some(new_node);
        self.size += 1;
    }

    /// 删除指定位置的节点
    fn delete_at_index(&mut self, index: i32) {
        if index < 0 || index >= self.size {
            return;
        }

        if index == 0 {
            self.head = self.head.take().unwrap().next;
        } else {
            let mut curr = self.head.as_mut().unwrap();
            for _ in 0..index - 1 {
                curr = curr.next.as_mut().unwrap();
            }
            curr.next = curr.next.as_mut().unwrap().next.take();
        }
        self.size -= 1;
    }
}

fn main() {
    let mut linked_list = MyLinkedList::new();

    // 测试用例
    println!("测试 addAtHead:");
    linked_list.add_at_head(1);
    println!("链表: {:?}", linked_list.head);

    println!("\n测试 addAtTail:");
    linked_list.add_at_tail(3);
    println!("链表: {:?}", linked_list.head);

    println!("\n测试 addAtIndex:");
    linked_list.add_at_index(1, 2);
    println!("链表: {:?}", linked_list.head);

    println!("\n测试 get:");
    println!("获取索引1的值: {}", linked_list.get(1));

    println!("\n测试 deleteAtIndex:");
    linked_list.delete_at_index(1);
    println!("链表: {:?}", linked_list.head);

    println!("\n测试 get:");
    println!("获取索引1的值: {}", linked_list.get(1));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_linked_list() {
        let mut linked_list = MyLinkedList::new();

        // 测试 addAtHead
        linked_list.add_at_head(1);
        assert_eq!(linked_list.get(0), 1);

        // 测试 addAtTail
        linked_list.add_at_tail(3);
        assert_eq!(linked_list.get(1), 3);

        // 测试 addAtIndex
        linked_list.add_at_index(1, 2);
        assert_eq!(linked_list.get(1), 2);

        // 测试 deleteAtIndex
        linked_list.delete_at_index(1);
        assert_eq!(linked_list.get(1), 3);

        // 测试边界情况
        assert_eq!(linked_list.get(-1), -1);
        assert_eq!(linked_list.get(2), -1);
    }
}
