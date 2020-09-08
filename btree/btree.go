package btree

import "fmt"

type Interface interface {
	Add(node *TreeNode)
	Delete(key string) *TreeNode
	Get(key string)
	Update(node *TreeNode)
}


type NodeType int
const (
	LEAF NodeType = 1 + iota
	UNLEAF
	ROOT
)

const MAX_SIZE = 16

type NodeValue struct {
	version int
	next *NodeValue
	data string
}

type TreeNode struct {
	size int
	nodeType NodeType

	key   string
	value NodeValue

	parent *TreeNode

	children []TreeNode
}

func (t *TreeNode) String() string{
	return fmt.Sprintf("key:%s value:%s", t.key, t.value.data)
}

func (t *TreeNode) Children() []TreeNode{
	return t.children
}


type BTree struct {
	root *TreeNode
}


func (t *BTree) Add(node *TreeNode) {
	root := t.root

	parent, position := getInsertPosition(root, node.key)

	if isFullNode(parent) {
		fmt.Println("full")
	}
	insertNode(parent, position, node)
	parent.size ++
	fmt.Println(parent.size)
}

func (t *BTree) Get(key string) *TreeNode{
	if t == nil {
		return nil
	}

	currentNode := t.root

	if currentNode.children == nil  {
		return nil
	}

	for currentNode.children != nil {
		index, eq := binarySearch(currentNode.children, key)
		node := &currentNode.children[index]

		if currentNode.nodeType == LEAF {
			if eq {
				return node
			} else {
				return nil
			}
		}
		currentNode = node
	}

	return nil
}

func (t *BTree) GetRoot() *TreeNode {
	return t.root
}

func (t *BTree) releaseNode(node *TreeNode) {
	p := node.parent
	for idx, child := range p.children {
		if &child == node {
			removeNode(node, idx)
			break
		}
	}
}

func NewBTree() (*BTree){
	return &BTree{
		root: NewTreeNode("root", "root", 0),
	}
}

func NewTreeNode(key string, data string, transactionID int) (*TreeNode) {
	return &TreeNode{
		size: 0,
		nodeType: LEAF,
		key: key,
		value: NodeValue{
			version: transactionID,
			data: data,
			next:nil,
		},
		children: nil,
	}
}

// index 左包含右不包含 [left, right)
// 返回节点插入位置
func binarySearch(index []TreeNode, target string) (int, bool) {
	if index == nil {
		return 0, false
	}
	start := 0
	end := len(index)
	var mid int

	if index[start].key > target {
		return -1, false
	}
	if index[len(index) -1].key < target {
		return len(index) - 1, false
	}

	for true {
		mid = (start + end) / 2
		if index[mid].key == target{
			return mid - 1, true
		}

		if start >= end {
			return start, false
		}

		if index[mid].key > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0, false
}

func getInsertPosition(node *TreeNode, key string) (parent *TreeNode, position int){

	currentNode := node

	for true {
		if currentNode.nodeType == LEAF {
			index, _ := binarySearch(currentNode.children, key)
			return currentNode, index
		} else {
			index, _ := binarySearch(currentNode.children, key)
			currentNode = & currentNode.children[index]
		}
	}

	return currentNode, 0
}

func isFullNode(node *TreeNode) bool {
	return len(node.children) >= MAX_SIZE
}

func insertNode(parent *TreeNode, position int, val *TreeNode) {
	index := parent.children
	last := append([]TreeNode{}, index[position:]...)
	index = append(index[0:position], *val)
	index = append(index, last...)

	parent.children = index
}

func removeNode(parent *TreeNode, position int) {
	index := parent.children
	if len(index) <= position {
		return
	}

	index = append([]TreeNode{}, index[0: position-1]...)
	index = append(index, index[position:]...)
	parent.children = index
}

func bisectionCutting(array []TreeNode) (left []TreeNode, right []TreeNode) {
	mid := len(array) / 2
	left = append([]TreeNode{}, array[0: mid]...)
	right = append([]TreeNode{}, array[mid: len(array) -1]...)
	return
}

