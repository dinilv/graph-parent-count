package main

import "fmt"

/*

Suppose we have some input data describing a graph of relationships between parents and children over multiple families and generations. The data is formatted as a list of (parent, child) pairs, where each individual is assigned a unique positive integer identifier.

For example, in this diagram, 3 is a child of 1 and 2, and 5 is a child of 4:

1   2    4           30
 \ /   /  \           \
  3   5    9  15      16
   \ / \    \ /
    6   7   12


Sample input/output (pseudodata):

pairs = [
    (5, 6), (1, 3), (2, 3), (3, 6), (15, 12),
    (5, 7), (4, 5), (4, 9), (9, 12), (30, 16)
]


Write a function that takes this data as input and returns two collections: one containing all individuals with zero known parents, and one containing all individuals with exactly one known parent.


Output may be in any order:

findNodesWithZeroAndOneParents(pairs) => [
  [1, 2, 4, 15, 30],   // Individuals with zero parents
  [5, 7, 9, 16]         // Individuals with exactly one parent
]

Complexity Analysis variables:

n: number of pairs in the input
*/

// model for keeping temperory metadata for graph node scan
type NodeData struct {
	Name        int
	ParentCount int8 //minimal data type for counter
}

// stretegy pattern with behaviourial utility
func (node *NodeData) IncrementParentCounter() {
	node.ParentCount++
}

func (node *NodeData) GetParentCounter() int8 {
	return node.ParentCount
}

// existence of node in the lookup complexity:O(1)
var nodeLookUp map[int]*NodeData

// node factory initialized
func NewNodeData(node int) *NodeData {
	return &NodeData{
		Name:        node,
		ParentCount: 0,
	}
}

func main() {
	pairs := [][]int{
		[]int{5, 6},
		[]int{1, 3},
		[]int{2, 3},
		[]int{3, 6},
		[]int{15, 12},
		[]int{5, 7},
		[]int{4, 5},
		[]int{4, 9},
		[]int{9, 12},
		[]int{30, 16},
	}
	// scanning through the 2D array
	ParseGraph(pairs)
	// iterate sructured output
	withoutParent := DisplayGraphWithCounter(0)
	fmt.Println("Without parent:", withoutParent)

	withoutOneParent := DisplayGraphWithCounter(1)
	fmt.Println("With one parent:", withoutOneParent)

}

// complexity:O(n)
func ParseGraph(pairs [][]int) ([]*NodeData, error) {
	// TODO:handle zero-length array & return error
	nodeDetails := []*NodeData{}
	// each scan results in different lookup
	nodeLookUp = make(map[int]*NodeData)
	for _, eachPair := range pairs {
		// format 0:parent, 1:child
		parent := eachPair[0]
		child := eachPair[1]
		// avoid existing node check
		node, isExist := nodeLookUp[child] // only child
		fmt.Println("node", node)
		if isExist {
			// increase the parent count
			node.IncrementParentCounter()
			nodeLookUp[child] = node
		} else {
			nodeNew := NewNodeData(child)
			nodeNew.IncrementParentCounter()
			nodeLookUp[child] = nodeNew
		}
		// check against parent node
		node, isExist = nodeLookUp[parent]
		if isExist {
			nodeLookUp[parent] = node
		} else {
			nodeNew := NewNodeData(parent)
			nodeLookUp[parent] = nodeNew
		}

	}

	fmt.Println("lookup", len(nodeLookUp))
	return nodeDetails, nil
}

// complexity:O(2n)
func DisplayGraphWithCounter(criteria int8) []int {
	fmt.Println("display")
	// iterate over the look up and display if criteria meets
	nodes := []int{}
	for key, node := range nodeLookUp {
		fmt.Println("node filter", node, key)
		// filter by criteria
		if node.GetParentCounter() == criteria {
			nodes = append(nodes, node.Name)
		}
	}
	return nodes
}
