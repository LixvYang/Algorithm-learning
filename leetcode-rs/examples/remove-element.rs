/*
 * 移除元素
 * https://leetcode.cn/problems/remove-element/
 *
 * 题目描述：
 * 给你一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，
 * 并返回移除后数组的新长度。
 *
 * 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
 *
 * 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 *
 * 示例 1:
 * 输入：nums = [3,2,2,3], val = 3
 * 输出：2, nums = [2,2]
 * 解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
 * 你不需要考虑数组中超出新长度后面的元素。
 *
 * 示例 2:
 * 输入：nums = [0,1,2,2,3,0,4,2], val = 2
 * 输出：5, nums = [0,1,3,0,4]
 * 解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
 * 注意这五个元素可为任意顺序。
 *
 * 提示：
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 50
 * 0 <= val <= 100
 */

/// 移除元素实现
///
/// # 参数
/// * `nums` - 需要修改的数组
/// * `val` - 需要移除的值
///
/// # 返回值
/// * 移除后数组的新长度
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(1)
fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut slow = 0;

    for fast in 0..nums.len() {
        if nums[fast] != val {
            nums[slow] = nums[fast];
            slow += 1;
        }
    }

    slow as i32
}

fn main() {
    // 测试用例
    let test_cases = vec![
        (vec![3, 2, 2, 3], 3, 2),
        (vec![0, 1, 2, 2, 3, 0, 4, 2], 2, 5),
        (vec![], 0, 0),
        (vec![1], 1, 0),
        (vec![1, 1, 1, 1], 1, 0),
        (vec![4, 5], 5, 1),
    ];

    println!("移除元素测试结果：");
    println!("-------------------");

    for (mut nums, val, expected) in test_cases {
        let original_nums = nums.clone();
        let result = remove_element(&mut nums, val);

        println!("输入: nums = {:?}, val = {}", original_nums, val);
        println!("输出: {}", result);
        println!("修改后的数组: {:?}", &nums[..result as usize]);
        println!("预期长度: {}", expected);
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
    fn test_remove_element() {
        let mut nums1 = vec![3, 2, 2, 3];
        assert_eq!(remove_element(&mut nums1, 3), 2);
        assert_eq!(&nums1[..2], &[2, 2]);

        let mut nums2 = vec![0, 1, 2, 2, 3, 0, 4, 2];
        assert_eq!(remove_element(&mut nums2, 2), 5);

        let mut nums3 = vec![];
        assert_eq!(remove_element(&mut nums3, 0), 0);

        let mut nums4 = vec![1];
        assert_eq!(remove_element(&mut nums4, 1), 0);
    }
}
