/*
 * 螺旋矩阵 II
 * https://leetcode.cn/problems/spiral-matrix-ii/
 *
 * 题目描述：
 * 给你一个正整数 n，生成一个包含 1 到 n^2 所有元素，
 * 且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix。
 *
 * 示例 1：
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 * 示例 2：
 * 输入：n = 1
 * 输出：[[1]]
 *
 * 提示：
 * 1 <= n <= 20
 */

/// 生成螺旋矩阵
///
/// # 参数
/// * `n` - 矩阵的大小
///
/// # 返回值
/// * n x n 的螺旋矩阵
///
/// # 时间复杂度
/// * O(n^2)
///
/// # 空间复杂度
/// * O(n^2)
fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
    let n = n as usize;
    let mut matrix = vec![vec![0; n]; n];
    let mut start_x = 0;
    let mut start_y = 0;
    let mut loop_count = n / 2;
    let mid = n / 2;
    let mut count = 1;
    let mut offset = 1;

    while loop_count > 0 {
        let mut i = start_x;
        let mut j = start_y;

        // 从左到右
        while j < n - offset {
            matrix[i][j] = count;
            count += 1;
            j += 1;
        }

        // 从上到下
        while i < n - offset {
            matrix[i][j] = count;
            count += 1;
            i += 1;
        }

        // 从右到左
        while j > start_y {
            matrix[i][j] = count;
            count += 1;
            j -= 1;
        }

        // 从下到上
        while i > start_x {
            matrix[i][j] = count;
            count += 1;
            i -= 1;
        }

        start_x += 1;
        start_y += 1;
        offset += 1;
        loop_count -= 1;
    }

    // 如果n是奇数，填充中心点
    if n % 2 == 1 {
        matrix[mid][mid] = count;
    }

    matrix
}

fn main() {
    // 测试用例
    let test_cases = vec![
        (1, vec![vec![1]]),
        (2, vec![vec![1, 2], vec![4, 3]]),
        (3, vec![vec![1, 2, 3], vec![8, 9, 4], vec![7, 6, 5]]),
        (
            4,
            vec![
                vec![1, 2, 3, 4],
                vec![12, 13, 14, 5],
                vec![11, 16, 15, 6],
                vec![10, 9, 8, 7],
            ],
        ),
    ];

    println!("螺旋矩阵 II 测试结果：");
    println!("-------------------");

    for (n, expected) in test_cases {
        let result = generate_matrix(n);

        println!("输入: n = {}", n);
        println!("输出:");
        for row in &result {
            println!("{:?}", row);
        }
        println!("预期:");
        for row in &expected {
            println!("{:?}", row);
        }
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
    fn test_generate_matrix() {
        assert_eq!(generate_matrix(1), vec![vec![1]]);
        assert_eq!(generate_matrix(2), vec![vec![1, 2], vec![4, 3]]);
        assert_eq!(
            generate_matrix(3),
            vec![vec![1, 2, 3], vec![8, 9, 4], vec![7, 6, 5]]
        );
        assert_eq!(
            generate_matrix(4),
            vec![
                vec![1, 2, 3, 4],
                vec![12, 13, 14, 5],
                vec![11, 16, 15, 6],
                vec![10, 9, 8, 7]
            ]
        );
    }
}
