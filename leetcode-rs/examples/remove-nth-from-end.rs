/*
 * 删除链表的倒数第 N 个结点
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 *
 * 题目描述：
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 * 示例 1：
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 * 示例 2：
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 * 示例 3：
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 * 提示：
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
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

/// 删除链表的倒数第 N 个结点实现（快慢指针法）
///
/// # 参数
/// * `head` - 链表的头节点
/// * `n` - 要删除的倒数第 n 个节点
///
/// # 返回值
/// * 删除节点后的链表头节点
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(1)
fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
    // 创建虚拟头节点
    let mut dummy = Box::new(ListNode::new(0));
    dummy.next = head;
    let mut fast = &dummy.clone();
    let mut slow = dummy.as_mut();

    // 快指针先移动 n 步
    for _ in 0..n {
        fast = fast.next.as_ref().unwrap();
    }

    // 快慢指针一起移动，直到快指针到达末尾
    while fast.next.is_some() {
        fast = fast.next.as_ref().unwrap();
        slow = slow.next.as_mut().unwrap();
    }

    // 删除倒数第 n 个节点
    let next = slow.next.as_mut().unwrap().next.take();
    slow.next = next;

    dummy.next
}

// 辅助函数：从数组创建链表
fn create_list(nums: &[i32]) -> Option<Box<ListNode>> {
    if nums.is_empty() {
        return None;
    }

    let mut head = Box::new(ListNode::new(nums[0]));
    let mut curr = &mut head;

    for &num in &nums[1..] {
        curr.next = Some(Box::new(ListNode::new(num)));
        curr = curr.next.as_mut().unwrap();
    }

    Some(head)
}

// 辅助函数：将链表转换为数组
fn list_to_vec(head: Option<Box<ListNode>>) -> Vec<i32> {
    let mut result = Vec::new();
    let mut curr = head;

    while let Some(node) = curr {
        result.push(node.val);
        curr = node.next;
    }

    result
}

fn main() {
    // 测试用例
    let test_cases = vec![
        (vec![1, 2, 3, 4, 5], 2, vec![1, 2, 3, 5]),
        (vec![1], 1, vec![]),
        (vec![1, 2], 1, vec![1]),
        (vec![1, 2, 3], 3, vec![2, 3]),
    ];

    println!("删除链表的倒数第 N 个结点测试结果：");
    println!("-------------------");

    for (nums, n, expected) in test_cases {
        let head = create_list(&nums);
        let result = remove_nth_from_end(head, n);
        let result_vec = list_to_vec(result);

        println!("输入: nums = {:?}, n = {}", nums, n);
        println!("输出: {:?}", result_vec);
        println!("预期: {:?}", expected);
        println!(
            "结果: {}",
            if result_vec == expected {
                "通过"
            } else {
                "失败"
            }
        );
        println!("-------------------");
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_remove_nth_from_end() {
        assert_eq!(
            list_to_vec(remove_nth_from_end(create_list(&[1, 2, 3, 4, 5]), 2)),
            vec![1, 2, 3, 5]
        );
        assert_eq!(
            list_to_vec(remove_nth_from_end(create_list(&[1]), 1)),
            vec![]
        );
        assert_eq!(
            list_to_vec(remove_nth_from_end(create_list(&[1, 2]), 1)),
            vec![1]
        );
        assert_eq!(
            list_to_vec(remove_nth_from_end(create_list(&[1, 2, 3]), 3)),
            vec![2, 3]
        );
    }
}
