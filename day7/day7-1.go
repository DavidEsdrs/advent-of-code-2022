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

func (node *Node) addChild(child ...*Node) {
	node.children = append(node.children, child...)
	for _, c := range child {
		node.size += c.size
	}
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
			exec(pieces[1:], &stack)
		case "dir":
			createFolder(pieces[1:], &stack)
		default:
			createFile(pieces, &stack)
		}
		lines++
	}

	println(subWithLessThan100k, calcSize(&root), subWithLessThan100k)
}

func exec(str []string, stack *Stack) error {
	switch str[0] {
	case "cd":
		current := stack.read()
		if str[1] == ".." {
			stack.pop()
			tabs = tabs[2:]
			return nil
		}
		node, found := searchChild(current.children, str[1], stack)
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

func searchChild(children []*Node, name string, stack *Stack) (node *Node, found bool) {
	for _, child := range children {
		if child.name == name {
			return child, true
		}
	}
	return nil, false
}

func createFolder(out []string, stack *Stack) {
	name := out[0]
	node := Node{name: name}
	stack.read().addChild(&node)
}

func createFile(out []string, stack *Stack) {
	sizeAsStr := out[0]
	size, err := strconv.Atoi(sizeAsStr)
	if err != nil {
		panic("Can't convert")
	}
	name := out[1]
	node := Node{name: name, size: uint64(size)}

	stack.read().addChild(&node)
}

func calcSize(root *Node) uint64 {
	var res uint64 = 0

	for _, child := range root.children {
		if len(child.children) > 0 {
			res += calcSize(child)
		} else {
			res += child.size
		}
	}

	if res <= 100000 {
		subWithLessThan100k += res
	}

	return res
}
