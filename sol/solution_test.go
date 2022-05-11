package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	temp := make([]*TreeNode, len(*input))
	for idx, val := range *input {
		num := 0
		if val != "null" {
			num, _ = strconv.Atoi(val)
		}
		if num != 0 {
			temp[idx] = &TreeNode{Val: num}
		}
	}
	for idx, val := range temp {
		if val != nil {
			if idx == 0 {
				tree = val
			}
			if 2*idx+1 < len(temp) {
				val.Left = temp[2*idx+1]
			}
			if 2*idx+2 < len(temp) {
				val.Right = temp[2*idx+2]
			}
		}
	}
	return tree
}
func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root = [3,9,20,null,null,15,7]",
			args: args{CreateBinaryTree(&[]string{"3", "9", "20", "null", "null", "15", "7"})},
			want: true,
		},
		{
			name: "root = [1,2,2,3,3,null,null,4,4]",
			args: args{CreateBinaryTree(&[]string{"1", "2", "2", "3", "3", "null", "null", "4", "4"})},
			want: false,
		},
		{
			name: "root = []",
			args: args{CreateBinaryTree(&[]string{})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
