/*
 * 移除链表元素
 * https://leetcode.cn/problems/remove-linked-list-elements/
 *
 * 题目描述：
 * 给你一个链表的头节点 head 和一个整数 val，
 * 请你删除链表中所有满足 Node.val == val 的节点，并返回新的头节点。
 *
 * 示例 1：
 * 输入：head = [1,2,6,3,4,5,6], val = 6
 * 输出：[1,2,3,4,5]
 *
 * 示例 2：
 * 输入：head = [], val = 1
 * 输出：[]
 *
 * 示例 3：
 * 输入：head = [7,7,7,7], val = 7
 * 输出：[]
 *
 * 提示：
 * 列表中的节点数目在范围 [0, 10^4] 内
 * 1 <= Node.val <= 50
 * 0 <= val <= 50
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

/// 移除链表元素实现
///
/// # 参数
/// * `head` - 链表的头节点
/// * `val` - 要移除的值
///
/// # 返回值
/// * 新的头节点
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(1)
fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    let mut dummy_head = Box::new(ListNode::new(0));
    dummy_head.next = head;
    let mut pre = dummy_head.as_mut();

    while let Some(cur) = pre.next.take() {
        if cur.val == val {
            pre.next = cur.next;
        } else {
            pre.next = Some(cur);
            pre = pre.next.as_mut().unwrap();
        }
    }
    dummy_head.next
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
        (vec![1, 2, 6, 3, 4, 5, 6], 6, vec![1, 2, 3, 4, 5]),
        (vec![], 1, vec![]),
        (vec![7, 7, 7, 7], 7, vec![]),
        (vec![1, 2, 2, 1], 2, vec![1, 1]),
    ];

    println!("移除链表元素测试结果：");
    println!("-------------------");

    for (nums, val, expected) in test_cases {
        let head = create_list(&nums);
        let result = remove_elements(head, val);
        let result_vec = list_to_vec(result);

        println!("输入: head = {:?}, val = {}", nums, val);
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

        assert_eq!(result_vec, expected);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_remove_elements() {
        assert_eq!(
            list_to_vec(remove_elements(create_list(&[1, 2, 6, 3, 4, 5, 6]), 6)),
            vec![1, 2, 3, 4, 5]
        );
        assert_eq!(list_to_vec(remove_elements(create_list(&[]), 1)), vec![]);
        assert_eq!(
            list_to_vec(remove_elements(create_list(&[7, 7, 7, 7]), 7)),
            vec![]
        );
        assert_eq!(
            list_to_vec(remove_elements(create_list(&[1, 2, 2, 1]), 2)),
            vec![1, 1]
        );
    }
}
