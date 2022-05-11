# golang_balanced_binary_tree

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
> 

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: true

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

```

**Example 3:**

```
Input: root = []
Output: true

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5000]`.
- $`-10^4$ <= Node.val <= $10^4$`

## 解析

給一個 二元樹的根結點 root 要判斷，該樹是不是 Balanced

定義一個 tree 是 balance tree 代表對 tree 所有結點的左右子結點高度差距不超過 1

從這個定義可以發現

假設一個 balanced tree 的根結點 root 必須會有下面條件

1. root 是空的 因為沒有子結點，所以是 balanced
2. root 非空，則 root.Left 的最大高度 與 root.Right 最大高度相差不超過1
3. root.Left 形成的 tree 必須是 Balance, root.Right 形成的 tree 必須是 Balance

## 程式碼

```go
package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return (Abs(MaxDepth(root.Left)-MaxDepth(root.Right)) <= 1) && isBalanced(root.Left) && isBalanced(root.Right)
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return Max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
## 困難點

1. 要找出 Balanced Tree的條件
2. 知道如何檢驗 balanced tree

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity