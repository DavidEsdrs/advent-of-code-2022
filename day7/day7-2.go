package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stack struct {
	_stack []*Node
}

func (stack *Stack) read() *Node {
	last := stack._stack[len(stack._stack)-1]
	return last
}

func (stack *Stack) pop() *Node {
	last := stack._stack[len(stack._stack)-1]
	stack._stack = stack._stack[:len(stack._stack)-1]
	return last
}

func (stack *Stack) push(node *Node) int {
	stack._stack = append(stack._stack, node)
	return len(stack._stack)
}

type Node struct {
	name     string
	children []*Node
	size     uint64
}

func (node *Node) AddChild(child ...*Node) {
	node.children = append(node.children, child...)
	for _, c := range child {
		node.size += c.size
	}
}

func (node *Node) IsDir() bool {
	return len(node.children) > 0
}

func (node *Node) GetSize() uint64 {
	if node.IsDir() {
		var size uint64 = 0
		for _, child := range node.children {
			size += child.GetSize()
		}
		return size
	}

	return node.size
}

var subWithLessThan100k uint64 = 0
var lines int = 0
var delay time.Duration = 0
var tabs string = "--"

func main() {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatal("Can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := Node{name: "/"}

	stack := Stack{}

	stack.push(&root)

	fmt.Printf(root.name + "\n")

	for scanner.Scan() {
		line := scanner.Text()

		pieces := strings.Split(line, " ")

		switch pieces[0] {
		case "$":
			Exec(pieces[1:], &stack)
		case "dir":
			CreateFolder(pieces[1:], &stack)
		default:
			CreateFile(pieces, &stack)
		}
		lines++
	}

	totalAvaiable := 70_000_000
	neededSpace := 30_000_000
	totalSize := root.GetSize()
	totalAvaiable -= int(totalSize)
	neededSpace -= totalAvaiable

	minSize := totalSize

	print("total_avaiable=", totalAvaiable, ", needed_space=", neededSpace, ", total_size=", totalSize, "\n")
	println("size=", calcMinimunSize(&root, uint64(neededSpace), &minSize))
	println("min=", minSize)
}

func Exec(str []string, stack *Stack) error {
	switch str[0] {
	case "cd":
		current := stack.read()
		if str[1] == ".." {
			stack.pop()
			tabs = tabs[2:]
			return nil
		}
		node, found := SearchChild(current.children, str[1], stack)
		if found {
			tabs += "--"
			stack.push(node)
			fmt.Printf("%v %v\n", tabs, node.name)
			return nil
		}
	case "ls":
		return nil
	}
	return nil
}

func SearchChild(children []*Node, name string, stack *Stack) (node *Node, found bool) {
	for _, child := range children {
		if child.name == name {
			return child, true
		}
	}
	return nil, false
}

func CreateFolder(out []string, stack *Stack) {
	name := out[0]
	node := Node{name: name}
	stack.read().AddChild(&node)
}

func CreateFile(out []string, stack *Stack) {
	sizeAsStr := out[0]
	size, err := strconv.Atoi(sizeAsStr)
	if err != nil {
		panic("Can't convert")
	}
	name := out[1]
	node := Node{name: name, size: uint64(size)}

	stack.read().AddChild(&node)
}

func calcMinimunSize(root *Node, neededSpace uint64, minSize *uint64) uint64 {
	var res uint64 = 0

	for _, child := range root.children {
		size := child.GetSize()
		if child.IsDir() {
			size := calcMinimunSize(child, neededSpace, minSize)
			if size >= neededSpace && size < *minSize {
				*minSize = size
			}
		}
		res += size
	}

	return res
}
