/*
 * 反转链表
 * https://leetcode.cn/problems/reverse-linked-list/
 *
 * 题目描述：
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 *
 * 示例 1：
 * 输入：head = [1,2,3,4,5]
 * 输出：[5,4,3,2,1]
 *
 * 示例 2：
 * 输入：head = [1,2]
 * 输出：[2,1]
 *
 * 示例 3：
 * 输入：head = []
 * 输出：[]
 *
 * 提示：
 * 链表中节点的数目范围是 [0, 5000]
 * -5000 <= Node.val <= 5000
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

/// 反转链表实现（迭代方法）
///
/// # 参数
/// * `head` - 链表的头节点
///
/// # 返回值
/// * 反转后的链表头节点
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(1)
fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut prev = None;
    let mut curr = head;

    while let Some(mut node) = curr {
        let next = node.next;
        node.next = prev;
        prev = Some(node);
        curr = next;
    }

    prev
}

/// 反转链表实现（递归方法）
fn reverse_list_recursive(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    fn reverse_helper(
        curr: Option<Box<ListNode>>,
        prev: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match curr {
            None => prev,
            Some(mut node) => {
                let next = node.next;
                node.next = prev;
                reverse_helper(next, Some(node))
            }
        }
    }

    reverse_helper(head, None)
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
        (vec![1, 2, 3, 4, 5], vec![5, 4, 3, 2, 1]),
        (vec![1, 2], vec![2, 1]),
        (vec![], vec![]),
        (vec![1], vec![1]),
    ];

    println!("反转链表测试结果：");
    println!("-------------------");

    for (nums, expected) in test_cases {
        // 测试迭代方法
        let head = create_list(&nums);
        let result = reverse_list(head);
        let result_vec = list_to_vec(result);

        println!("输入: {:?}", nums);
        println!("迭代方法输出: {:?}", result_vec);
        println!("预期: {:?}", expected);
        println!(
            "结果: {}",
            if result_vec == expected {
                "通过"
            } else {
                "失败"
            }
        );

        // 测试递归方法
        let head = create_list(&nums);
        let result = reverse_list_recursive(head);
        let result_vec = list_to_vec(result);

        println!("递归方法输出: {:?}", result_vec);
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
    fn test_reverse_list() {
        // 测试迭代方法
        assert_eq!(
            list_to_vec(reverse_list(create_list(&[1, 2, 3, 4, 5]))),
            vec![5, 4, 3, 2, 1]
        );
        assert_eq!(list_to_vec(reverse_list(create_list(&[1, 2]))), vec![2, 1]);
        assert_eq!(list_to_vec(reverse_list(create_list(&[]))), vec![]);

        // 测试递归方法
        assert_eq!(
            list_to_vec(reverse_list_recursive(create_list(&[1, 2, 3, 4, 5]))),
            vec![5, 4, 3, 2, 1]
        );
        assert_eq!(
            list_to_vec(reverse_list_recursive(create_list(&[1, 2]))),
            vec![2, 1]
        );
        assert_eq!(
            list_to_vec(reverse_list_recursive(create_list(&[]))),
            vec![]
        );
    }
}
