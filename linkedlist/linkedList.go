package linkedlist

import (
	"errors"
	"fmt"
	"strconv"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value    int64
	NextNode *Node
}

type Iterator struct {
	CurrentNode Node
}

func Make() LinkedList {
	return LinkedList{}
}

func (list *LinkedList) Push(value int64) {
	if list.Head == nil {
		list.Head = &Node{
			Value:    value,
			NextNode: nil,
		}
	} else {
		oldInitNode := *list.Head
		*list.Head = Node{
			Value:    value,
			NextNode: &oldInitNode,
		}
	}
}

func (list *LinkedList) Put(after int64, value int64) {
	putAfter := list.Head
	if putAfter == nil {
		list.Head = &Node{
			Value:    value,
			NextNode: nil,
		}
	} else {
		for putAfter.Value != after {
			putAfter = putAfter.NextNode
		}
		tail := putAfter.NextNode
		putAfter.NextNode = &Node{
			Value:    value,
			NextNode: tail,
		}
	}
}

func (list *LinkedList) Append(value int64) {
	lastNode := list.Head
	if lastNode == nil {
		list.Head = &Node{
			Value:    value,
			NextNode: nil,
		}
	} else {
		for lastNode.NextNode != nil {
			lastNode = lastNode.NextNode
		}
		lastNode.NextNode = &Node{
			Value: value,
		}
	}
}

func (list *LinkedList) Pop() int64 {
	result := list.Head.Value
	list.Head = list.Head.NextNode
	return result
}

func (list *LinkedList) Remove(value int64) bool {
	current := list.Head
	prev := list.Head

	if current.Value == value {
		list.Head = current.NextNode
		return true
	}

	for current.NextNode != nil {
		prev = current
		current = current.NextNode
	}

	if current == nil {
		return false
	}
	prev.NextNode = current.NextNode
	return true
}

func (list *LinkedList) Get(index int64) (int64, error) {
	iterator := Iterator{CurrentNode: *list.Head}
	i := int64(0)
	for ; iterator.HasNext(); iterator.Next() {
		nextNodeValue := iterator.CurrentNode.Value
		if i == index {
			return nextNodeValue, nil
		}
		i++
	}
	if i == index {
		return iterator.CurrentNode.Value, nil
	}
	return -1, errors.New("No such element")
}

func (list *LinkedList) Belongs(value int64) bool {
	i := Iterator{CurrentNode: *list.Head}
	for ; i.HasNext(); i.Next() {
		nextNodeValue := i.CurrentNode.Value
		if nextNodeValue == value {
			return true
		}
	}
	return i.CurrentNode.Value == value
}

func (list *LinkedList) Print() {
	i := Iterator{CurrentNode: *list.Head}
	for ; i.HasNext(); i.Next() {
		fmt.Print(strconv.FormatInt(i.CurrentNode.Value, 10) + "->")
	}
	if &i.CurrentNode != nil {
		fmt.Print(strconv.FormatInt(i.CurrentNode.Value, 10) + "\n")
	}
}

func (i *Iterator) HasNext() bool {
	return !(i.CurrentNode.NextNode == nil)
}

func (i *Iterator) Next() Node {
	i.CurrentNode = *i.CurrentNode.NextNode
	return i.CurrentNode
}
