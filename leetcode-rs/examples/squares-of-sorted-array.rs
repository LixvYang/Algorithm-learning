/*
 * 有序数组的平方
 * https://leetcode.cn/problems/squares-of-a-sorted-array/
 *
 * 题目描述：
 * 给你一个按非递减顺序排序的整数数组 nums，返回每个数字的平方组成的新数组，
 * 要求也按非递减顺序排序。
 *
 * 示例 1：
 * 输入：nums = [-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 解释：平方后，数组变为 [16,1,0,9,100]
 * 排序后，数组变为 [0,1,9,16,100]
 *
 * 示例 2：
 * 输入：nums = [-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 * 提示：
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 已按非递减顺序排序
 */

/// 有序数组的平方实现
///
/// # 参数
/// * `nums` - 非递减顺序排序的整数数组
///
/// # 返回值
/// * 平方后的有序数组
///
/// # 时间复杂度
/// * O(n)
///
/// # 空间复杂度
/// * O(n)

pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
    let (mut i, mut j, mut k) = (0, nums.len() - 1, nums.len() - 1);
    let mut result = vec![0; nums.len()];
    while i <= j {
        let (lm, rm) = (
            nums[i as usize] * nums[i as usize],
            nums[j as usize] * nums[j as usize],
        );

        if lm < rm {
            result[k] = rm;
            j = j - 1;
        } else {
            result[k] = lm;
            i = i + 1;
        }
        k = k - 1;
    }
    result
}

fn main() {
    // 测试用例
    let test_cases = vec![
        (vec![-4, -1, 0, 3, 10], vec![0, 1, 9, 16, 100]),
        // (vec![-7, -3, 2, 3, 11], vec![4, 9, 9, 49, 121]),
        // (vec![0], vec![0]),
        // (vec![-5, -3, -2, -1], vec![1, 4, 9, 25]),
        // (vec![1, 2, 3, 4, 5], vec![1, 4, 9, 16, 25]),
    ];

    println!("有序数组的平方测试结果：");
    println!("-------------------");

    for (nums, expected) in test_cases {
        let result = sorted_squares(nums.clone());

        println!("输入: nums = {:?}", nums);
        println!("输出: {:?}", result);
        println!("预期: {:?}", expected);
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
    fn test_sorted_squares() {
        assert_eq!(
            sorted_squares(vec![-4, -1, 0, 3, 10]),
            vec![0, 1, 9, 16, 100]
        );
        assert_eq!(
            sorted_squares(vec![-7, -3, 2, 3, 11]),
            vec![4, 9, 9, 49, 121]
        );
        assert_eq!(sorted_squares(vec![0]), vec![0]);
        assert_eq!(sorted_squares(vec![-5, -3, -2, -1]), vec![1, 4, 9, 25]);
        assert_eq!(sorted_squares(vec![1, 2, 3, 4, 5]), vec![1, 4, 9, 16, 25]);
    }
}
