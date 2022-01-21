// Package leetcode provides 
package leetcode

// 递归法后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode)  {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

// 迭代法后序遍历
func postorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)

		if node.Left != nil {
			st.PushBack(node.Left)
		}

		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}

	reverse(ans)

	return ans
}

func reverse(ans []int) {
	l, r := 0, len(ans)-1
	for l<r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}