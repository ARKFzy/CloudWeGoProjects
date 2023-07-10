package main

import "fmt"

type LinkNode struct {
	value int
	next  *LinkNode
}

func NewLinkNode(value int) *LinkNode {
	return &LinkNode{
		value: value,
		next:  nil,
	}
}

func (n *LinkNode) GetValue() int {
	return n.value
}

func (n *LinkNode) SetNext(next *LinkNode) {
	n.next = next
}
func main() {
	link1 := NewLinkNode(1)
	link1.next = NewLinkNode(888)
	for v := link1; ; v = v.next {
		fmt.Println(v.GetValue())
		if v.next == nil {
			break
		}
	}

}
