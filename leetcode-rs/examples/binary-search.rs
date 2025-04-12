/*
 * 二分查找
 * https://leetcode.cn/problems/binary-search/
 *
 * 题目描述：
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
 * 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
 *
 * 示例 1:
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 *
 * 示例 2:
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 *
 * 提示：
 * 1. 你可以假设 nums 中的所有元素是不重复的。
 * 2. n 将在 [1, 10000]之间。
 * 3. nums 的每个元素都将在 [-9999, 9999]之间。
 */

/// 二分查找实现
///
/// # 参数
/// * `nums` - 升序排列的整数数组
/// * `target` - 要查找的目标值
///
/// # 返回值
/// * 如果找到目标值，返回其索引
/// * 如果未找到目标值，返回 -1
///
/// # 时间复杂度
/// * O(log n)
///
/// # 空间复杂度
/// * O(1)
fn search(nums: Vec<i32>, target: i32) -> i32 {
    let mut left = 0;
    let mut right = nums.len() as i32 - 1;

    while left <= right {
        // 防止整数溢出
        let mid = left + (right - left) / 2;

        match nums[mid as usize].cmp(&target) {
            std::cmp::Ordering::Equal => return mid,
            std::cmp::Ordering::Less => left = mid + 1,
            std::cmp::Ordering::Greater => right = mid - 1,
        }
    }
    -1
}

fn main() {
    // 测试用例
    let test_cases = vec![
        (vec![-1, 0, 3, 5, 9, 12], 9, 4),  // 目标值在数组中
        (vec![-1, 0, 3, 5, 9, 12], 2, -1), // 目标值不在数组中
        (vec![5], 5, 0),                   // 只有一个元素且匹配
        (vec![5], -5, -1),                 // 只有一个元素且不匹配
        (vec![1, 2, 3, 4, 5], 1, 0),       // 目标值在数组开头
        (vec![1, 2, 3, 4, 5], 5, 4),       // 目标值在数组末尾
    ];

    println!("二分查找测试结果：");
    println!("-------------------");

    for (nums, target, expected) in test_cases {
        let result = search(nums.clone(), target);
        println!("输入: nums = {:?}, target = {}", nums, target);
        println!("输出: {}", result);
        println!("预期: {}", expected);
        println!(
            "结果: {}",
            if result == expected {
                "通过"
            } else {
                "失败"
            }
        );
        println!("-------------------");
        assert_eq!(result, expected);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_search() {
        assert_eq!(search(vec![-1, 0, 3, 5, 9, 12], 9), 4);
        assert_eq!(search(vec![-1, 0, 3, 5, 9, 12], 2), -1);
        assert_eq!(search(vec![5], 5), 0);
        assert_eq!(search(vec![5], -5), -1);
        assert_eq!(search(vec![1, 2, 3, 4, 5], 1), 0);
        assert_eq!(search(vec![1, 2, 3, 4, 5], 5), 4);
    }
}
