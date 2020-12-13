package main

import "testing"

func TestNil(t *testing.T) {
	var root *Tree
	root = nil
	got := sumMetadata(root)

	if got != 0 {
		t.Errorf("nil tree returned %d, want 0", got)
	}
}

func TestRootOnly(t *testing.T) {
	metadata := [2]int{12, 17}

	want := 0
	for _, v := range metadata {
		want += v
	}

	root := NewTree(metadata[:])
	got := sumMetadata(root)

	if got != want {
		t.Errorf("single root returned %d, want %d", got, want)
	}
}

func TestSimpleTree(t *testing.T) {
	metadata1 := [2]int{12, 17}
	metadata2 := [1]int{9}
	metadata3 := [3]int{4, 5, 6}

	want := 0
	for _, v := range metadata1 {
		want += v
	}
	for _, v := range metadata2 {
		want += v
	}
	for _, v := range metadata3 {
		want += v
	}

	
	root := NewTree(metadata1[:])
	root.Children = append(root.Children, NewTree(metadata2[:]))
	root.Children = append(root.Children, NewTree(metadata3[:]))
	
	got := sumMetadata(root)
	if got != want {
		t.Errorf("simple tree returned %d, want %d", got, want)
	}
}

func TestBigTree(t *testing.T) {
	metadata1 := [2]int{12, 17}
	metadata2 := [1]int{9}
	metadata3 := [3]int{4, 5, 6}
	metadata4 := [3]int{7, 8, 9}
	metadata5 := [3]int{21, 22, 23}

	want := 0
	for _, v := range metadata1 {
		want += v
	}
	for _, v := range metadata2 {
		want += v
	}
	for _, v := range metadata3 {
		want += v
	}
	for _, v := range metadata4 {
		want += v
	}
	for _, v := range metadata5 {
		want += v
	}

	grandchild1 := NewTree(metadata4[:])
	grandchild2 := NewTree(metadata5[:])

	child1 := NewTree(metadata2[:])
	child2 := NewTree(metadata3[:])

	AddChild(child1, grandchild1)
	AddChild(child2, grandchild2)
	
	root := NewTree(metadata1[:])

	AddChild(root, child1)
	AddChild(root, child2)
	
	got := sumMetadata(root)
	if got != want {
		t.Errorf("big tree returned %d, want %d", got, want)
	}
}

