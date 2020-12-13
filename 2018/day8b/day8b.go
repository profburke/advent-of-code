package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func read(r io.Reader) (data []int, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
		s := scanner.Text()
		var n int64
		n, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return
		}
		data = append(data, int(n))
	}
	
	err = scanner.Err()
	return data, err
}

func main() {
	data, err := read(os.Stdin)
	if err != nil {
		fmt.Println("trouble reading input: ", err)
		os.Exit(1)
	}

	tree := build(data)
	checksum := nodeValue(tree)
	fmt.Println(checksum)
}


func nodeValue(root *Tree) (value int) {
	if root == nil {
		value = 0
	} else if len(root.Children) == 0 {
		for _, v := range root.Metadata {
			value += v
		}
	} else {
		for _, i := range root.Metadata {
			index := i - 1
			if index < len(root.Children) {
				value += nodeValue(root.Children[index])
			}
		}
	}

	return value
}
