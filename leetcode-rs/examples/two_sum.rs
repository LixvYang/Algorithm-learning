//! 两数之和
//!
//! 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
//!
//! 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//!
//! 你可以按任意顺序返回答案。

use core::num;
use std::{collections::HashMap, hash::Hash, vec};

/// 使用哈希表实现的两数之和解法
///
/// # 时间复杂度
/// O(n)
///
/// # 空间复杂度
/// O(n)
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map = HashMap::new(); // value -> index

    for (i, &num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&j) = map.get(&complement) {
            return vec![j as i32, i as i32];
        }
        map.insert(num, i);
    }

    vec![]
}

pub fn two_sum_2(nums: Vec<i32>, target: i32) -> Vec<i32> {
    for i in 0..nums.len() {
        for j in i + 1..nums.len() {
            if nums[i] + nums[j] == target {
                return vec![i as i32, j as i32];
            }
        }
    }
    vec![]
}

fn main() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    let result = two_sum(nums.clone(), target);
    let result_2 = two_sum_2(nums.clone(), target);
    println!("Result: {:?}", result);
    println!("Result_2: {:?}", result_2);
}
