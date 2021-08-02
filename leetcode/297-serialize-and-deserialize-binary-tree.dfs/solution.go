package leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// serialize Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	builder := strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			builder.WriteString("null,")
			return
		}
		builder.WriteString(strconv.Itoa(node.Val))
		builder.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return builder.String()
}

// deserialize Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if vals[0] == "null" {
			vals = vals[1:]
			return nil
		}
		val, _ := strconv.Atoi(vals[0])
		vals = vals[1:]
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}
