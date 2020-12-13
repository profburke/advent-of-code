package main

func getMetadata(data []int, nMeta int, index int) (metadata []int) {
	for i := index; i < index + nMeta; i++ {
		metadata = append(metadata, data[i])
	}
	return metadata
}

func getChildren(data []int, nChildren int, start int) (children []*Tree, indexAfterChildren int) {
	index := start
	for i := 0; i < nChildren; i++ {
		var child *Tree
		child, index = builder(data, index)
		children = append(children, child)
	}

	indexAfterChildren = index
	return children, index
}

func builder(data []int, start int) (root *Tree, newStart int) {
	nChildren := data[start]
	nMeta := data[start + 1]

	children, metaIndex := getChildren(data, nChildren, start + 2)
	metadata := getMetadata(data, nMeta, metaIndex)

	root = NewTree(metadata)
	root.Children = children

	newStart = metaIndex + nMeta
	
	return root, newStart
}

func build(data []int) (root *Tree) {
	root, _ = builder(data, 0)
	return root
}
