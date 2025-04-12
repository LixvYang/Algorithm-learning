/*
 * 两两交换链表中的节点
 * https://leetcode.cn/problems/swap-nodes-in-pairs/
 *
 * 题目描述：
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
 * 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 *
 * 示例 1：
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 * 示例 2：
 * 输入：head = []
 * 输出：[]
 *
 * 示例 3：
 * 输入：head = [1]
 * 输出：[1]
 *
 * 提示：
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
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

/// 两两交换链表中的节点实现（类似 Go 的方式）
///
/// # 参数
/// * `head` - 链表的头节点
///
/// # 返回值
/// * 交换后的链表头节点
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(1)
fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut dummy = Box::new(ListNode::new(0));
    dummy.next = head;
    let mut cur = dummy.as_mut();

    while cur.next.is_some() && cur.next.as_ref().unwrap().next.is_some() {
        // 1. 保存第一个节点
        let mut first = cur.next.take();

        // 2. 获取第二个节点
        let mut second = first.as_mut().unwrap().next.take();

        // 3. 保存第三个节点开始的后续链表
        let rest = second.as_mut().unwrap().next.take();

        // 4. 重新连接节点
        // second->first->rest
        first.as_mut().unwrap().next = rest; // first->rest
        second.as_mut().unwrap().next = first; // second->first->rest
        cur.next = second; // cur->second->first->rest

        // 5. 移动 cur 指针到下一组的前一个位置
        cur = cur.next.as_mut().unwrap().next.as_mut().unwrap();
    }

    dummy.next
}

/// 递归方法实现两两交换
fn swap_pairs_recursive(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    // 如果链表为空或只有一个节点，直接返回
    if head.is_none() || head.as_ref().unwrap().next.is_none() {
        return head;
    }

    // 获取第一个节点
    let mut head = head.unwrap();
    // 获取第二个节点
    let mut next = head.next.take().unwrap();
    // 递归处理后续节点
    head.next = swap_pairs(next.next.take());
    // 交换节点
    next.next = Some(head);
    // 返回新的头节点（原来的第二个节点）
    Some(next)
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
        (vec![1, 2, 3, 4], vec![2, 1, 4, 3]),
        (vec![1], vec![1]),
        (vec![], vec![]),
        (vec![1, 2, 3], vec![2, 1, 3]),
    ];

    println!("两两交换链表中的节点测试结果：");
    println!("-------------------");

    for (nums, expected) in test_cases {
        let head = create_list(&nums);
        let result = swap_pairs(head);
        let result_vec = list_to_vec(result);

        println!("输入: {:?}", nums);
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
    fn test_swap_pairs() {
        assert_eq!(
            list_to_vec(swap_pairs(create_list(&[1, 2, 3, 4]))),
            vec![2, 1, 4, 3]
        );
        assert_eq!(list_to_vec(swap_pairs(create_list(&[1]))), vec![1]);
        assert_eq!(list_to_vec(swap_pairs(create_list(&[]))), vec![]);
        assert_eq!(
            list_to_vec(swap_pairs(create_list(&[1, 2, 3]))),
            vec![2, 1, 3]
        );
    }
}
