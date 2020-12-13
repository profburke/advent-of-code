package main

type Tree struct {
	Children []*Tree
	Metadata []int
}

func NewTree(metadata []int) *Tree {
	return &Tree{nil, metadata}
}

func AddChild(root *Tree, node *Tree) {
	if node == nil || root == nil {
		return
	}
	root.Children = append(root.Children, node)
}

func walk(root *Tree, visitor (func([]int) int)) (result int) {
	if root == nil {
		return 0
	}
	
	result = 0
	for _, child := range root.Children {
		result += walk(child, visitor)
	}
	result += visitor(root.Metadata)
	
	return result
}

func sumMetadata(root *Tree) int {
	metaSum := func(data []int) (sum int) {
		sum = 0
		for _, v := range data {
			sum += v
		}
		return sum
	}
	
	return walk(root, metaSum)
}
