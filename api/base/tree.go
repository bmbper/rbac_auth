package base

func BuildTree[T TreeNode](beans []T) []T {
	var roots []T
	nodeMap := make(map[string]T)

	// First pass: create a map of all nodes
	for _, bean := range beans {
		code := bean.GetCode()
		nodeMap[code] = bean
	}

	// Second pass: build parent-child relationships
	for _, bean := range beans {
		parentCode := bean.GetParentCode()
		if parentCode == "" || parentCode == TREE_NODE_ROOT {
			roots = append(roots, bean)
		} else if parent, exists := nodeMap[parentCode]; exists {
			children := parent.GetChildren()
			children = append(children, bean)
			parent.SetChildren(children)
		}
	}

	return roots
}
