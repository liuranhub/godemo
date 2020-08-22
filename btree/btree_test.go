package btree_test

import (
	"testing"
	"fmt"
	"godemo/btree"
)

func TestIndexInsert(t *testing.T) {
	var array []string

	for i:=0;i < 18 ; i ++  {
		array = append(array, fmt.Sprintf("%d", i))
	}

	tree := btree.NewBTree()

	for _, val := range array {
		tree.Add(btree.NewTreeNode(val, val, 1))
	}

	for _, val := range tree.GetRoot().Children(){
		fmt.Println(val.String())
	}

}

func TestBinarySearch(t *testing.T) {
	//array := []int{1,3,6,13,45,90,100}
	//
	//data := []int{-1,7,45,110}

	//for _, val := range data {
	//	index, success := btree.binarySearch(array, val)
	//
	//	fmt.Printf("position:%d success:%t\n", index, success)
	//}
}