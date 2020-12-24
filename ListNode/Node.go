package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}



func main() {
	root := CreateNode(3)
	root.Left = CreateNode(1)
	root.Left.Right = CreateNode(2)
	root.Right = &Node{5, nil, nil}
	root.Right.Left = CreateNode(4)

	fmt.Print("\n前序遍历: ")
	root.PreOrder()
	fmt.Print("\n非递归前序遍历: ")
	root.NonRecursivePreOrder()
	fmt.Print("\n中序遍历: ")
	root.MiddleOrder()
	fmt.Print("\n非递归中序遍历: ")
	root.NonRecursiveMiddleOrder()
	fmt.Print("\n后序遍历: ")
	root.PostOrder()
	fmt.Print("\n非递归后序遍历: ")
	root.NonRecursivePostOrder()
	fmt.Print("\n层次遍历: ")
	root.BreadthFirstSearch()
	//fmt.Println("\n层数: ", root.Layers())
	//fmt.Println("\n层数: ", root.LayersByQueue())}
}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

//先序遍历
func (root *Node) PreOrder() {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	root.Left.PreOrder()
	root.Right.PreOrder()
	//PreOderTravel(root.right)
}

//中序遍历
func (root *Node) MiddleOrder() {
	if root == nil {
		return
	}

	root.Left.MiddleOrder()
	fmt.Print(root.Value, " ")
	root.Right.MiddleOrder()
}

//后续遍历
func (root *Node) PostOrder(){
	if root == nil {
		return
	}
	root.Left.PostOrder()
	root.Right.PostOrder()
	fmt.Print(root.Value, " ")
}

//广度优先，层次遍历
func (root *Node) BreadthFirstSearch() {
	result := make([]int, 0)
	nodes := []*Node{root}

	for len(nodes) > 0 {
		curNode := nodes[0]
		result = append(result, curNode.Value)
		nodes = nodes[1:]
		if curNode.Left != nil {
			nodes = append(nodes, curNode.Left)
		}
		if curNode.Right != nil {
			nodes = append(nodes, curNode.Right)
		}
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//层数
func (root *Node) Layers()int {
	if root == nil {
		return 0
	}
	leftLayers  := root.Left.Layers()
	rightLayers := root.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

//前序非递归
func (root *Node) NonRecursivePreOrder() {
	if root == nil {
		return
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) > 0 {
		//先保存根节点，然后再遍历左子树
		//保存根节点，回来的时候遍历右子树
		for root != nil {
			result = append(result, root.Value)
			stack = append(stack, root)
			root = root.Left
		}
		//遍历右子树
		pop := stack[len(stack) - 1]
		root = pop.Right
		stack = stack[0:len(stack) - 1]
	}
	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//中序非递归

func (root *Node) NonRecursiveMiddleOrder() {
	if root == nil {
		return
	}

	result := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) > 0 {
		//先遍历左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		pop := stack[len(stack) - 1]
		result = append(result, pop.Value)
		root = pop.Right
		stack = stack[0:len(stack) - 1]
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//后序非递归
func (root *Node) NonRecursivePostOrder() {
	if root == nil {
		return
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)
	var lastVisit *Node
	for root != nil || len(stack) != 0 {
		//先遍历左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := stack[len(stack) - 1]
		//如果右节点为空则记录
		if last.Right == nil || last.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, last.Value)
			lastVisit = last
		} else {
			root = last.Right
		}
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}
}
